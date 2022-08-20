package codegen

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

type Resource struct {
	TableFunctionName string
	PackageName       string
	FileName          string
	ListFunction      string
	Struct            interface{}
	DefaultColumns    []codegen.ColumnDefinition
	Table             *codegen.TableDefinition
	Template          string
	Service           string
	StructField       string
	StructName        string
	Imports           []string
}

var ProjectIdColumn = codegen.ColumnDefinition{
	Name:     "project_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveProject",
}
