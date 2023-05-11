// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"sort"
	"text/template"
	"unicode"

	"code.gitea.io/gitea/modules/json"

	"golang.org/x/text/unicode/rangetable"
)

// ambiguous.json provides a one to one mapping of ambiguous characters to other characters
// See https://github.com/hediet/vscode-unicode-data/blob/main/out/ambiguous.json

type AmbiguousTable struct {
	Confusable []rune
	With       []rune
	Locale     string
	RangeTable *unicode.RangeTable
}

type RunePair struct {
	Confusable rune
	With       rune
}

var verbose bool

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `%s: Generate AmbiguousCharacter

Usage: %[1]s [-v] [-o output.go] ambiguous.json
`, os.Args[0])
		flag.PrintDefaults()
	}

	output := ""
	flag.BoolVar(&verbose, "v", false, "verbose output")
	flag.StringVar(&output, "o", "ambiguous_gen.go", "file to output to")
	flag.Parse()
	input := flag.Arg(0)
	if input == "" {
		input = "ambiguous.json"
	}

	bs, err := os.ReadFile(input)
	if err != nil {
		fatalf("Unable to read: %s Err: %v", input, err)
	}

	var unwrapped string
	if err := json.Unmarshal(bs, &unwrapped); err != nil {
		fatalf("Unable to unwrap content in: %s Err: %v", input, err)
	}

	fromJSON := map[string][]uint32{}
	if err := json.Unmarshal([]byte(unwrapped), &fromJSON); err != nil {
		fatalf("Unable to unmarshal content in: %s Err: %v", input, err)
	}

	tables := make([]*AmbiguousTable, 0, len(fromJSON))
	for locale, chars := range fromJSON {
		table := &AmbiguousTable{Locale: locale}
		table.Confusable = make([]rune, 0, len(chars)/2)
		table.With = make([]rune, 0, len(chars)/2)
		pairs := make([]RunePair, len(chars)/2)
		for i := 0; i < len(chars); i += 2 {
			pairs[i/2].Confusable, pairs[i/2].With = rune(chars[i]), rune(chars[i+1])
		}
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].Confusable < pairs[j].Confusable
		})
		for _, pair := range pairs {
			table.Confusable = append(table.Confusable, pair.Confusable)
			table.With = append(table.With, pair.With)
		}
		table.RangeTable = rangetable.New(table.Confusable...)
		tables = append(tables, table)
	}
	sort.Slice(tables, func(i, j int) bool {
		return tables[i].Locale < tables[j].Locale
	})
	data := map[string]interface{}{
		"Tables": tables,
	}

	if err := runTemplate(generatorTemplate, output, &data); err != nil {
		fatalf("Unable to run template: %v", err)
	}
}

func runTemplate(t *template.Template, filename string, data interface{}) error {
	buf := bytes.NewBuffer(nil)
	if err := t.Execute(buf, data); err != nil {
		return fmt.Errorf("unable to execute template: %w", err)
	}
	bs, err := format.Source(buf.Bytes())
	if err != nil {
		verbosef("Bad source:\n%s", buf.String())
		return fmt.Errorf("unable to format source: %w", err)
	}

	old, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to read old file %s because %w", filename, err)
	} else if err == nil {
		if bytes.Equal(bs, old) {
			// files are the same don't rewrite it.
			return nil
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s because %w", filename, err)
	}
	defer file.Close()
	_, err = file.Write(bs)
	if err != nil {
		return fmt.Errorf("unable to write generated source: %w", err)
	}
	return nil
}

var generatorTemplate = template.Must(template.New("ambiguousTemplate").Parse(`// This file is generated by modules/charset/ambiguous/generate.go DO NOT EDIT
// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT


package charset

import "unicode"

// This file is generated from https://github.com/hediet/vscode-unicode-data/blob/main/out/ambiguous.json

// AmbiguousTable matches a confusable rune with its partner for the Locale
type AmbiguousTable struct {
	Confusable []rune
	With       []rune
	Locale     string
	RangeTable *unicode.RangeTable
}

// AmbiguousCharacters provides a map by locale name to the confusable characters in that locale
var AmbiguousCharacters = map[string]*AmbiguousTable{
	{{range .Tables}}{{printf "%q:" .Locale}} {
			Confusable: []rune{ {{range .Confusable}}{{.}},{{end}} },
			With: []rune{ {{range .With}}{{.}},{{end}} },
			Locale: {{printf "%q" .Locale}},
			RangeTable: &unicode.RangeTable{
				R16: []unicode.Range16{
			{{range .RangeTable.R16 }}		{Lo:{{.Lo}}, Hi:{{.Hi}}, Stride: {{.Stride}}},
			{{end}}	},
				R32: []unicode.Range32{
			{{range .RangeTable.R32}}		{Lo:{{.Lo}}, Hi:{{.Hi}}, Stride: {{.Stride}}},
			{{end}}	},
				LatinOffset: {{.RangeTable.LatinOffset}},
			},
		},
	{{end}}
}

`))

func logf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

func verbosef(format string, args ...interface{}) {
	if verbose {
		logf(format, args...)
	}
}

func fatalf(format string, args ...interface{}) {
	logf("fatal: "+format+"\n", args...)
	os.Exit(1)
}
