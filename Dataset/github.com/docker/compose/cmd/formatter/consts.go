/*
   Copyright 2020 Docker Compose CLI authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package formatter

const (
	// JSON Print in JSON format
	JSON = "json"
	// TemplateLegacyJSON the legacy json formatting value using go template
	TemplateLegacyJSON = "{{json.}}"
	// PRETTY is the constant for default formats on list commands
	// Deprecated: use TABLE
	PRETTY = "pretty"
	// TABLE Print output in table format with column headers (default)
	TABLE = "table"
)
