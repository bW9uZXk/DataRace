package consts

const TemplatePbEntityMessageContent = `
// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

syntax = "proto3";

package {PackageName};

option go_package = "{GoPackage}";
{OptionContent}
{Imports}

{EntityMessage}
`
