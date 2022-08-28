package codegen

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

type Resource struct {
	// PackageName name is the packgename in the source plugin this resource is located
	PackageName string
	// DefaultColumns columns that will be appended to the main table
	DefaultColumns []codegen.ColumnDefinition
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// Struct that will be used to generate the cloudquery table
	Struct interface{}
	// StructName is the name of the Struct because it can't be inferred by reflection
	StructName string
	// Service is the name of the gcp service the struct/api is residing
	Service string
	// SubService is the name of the gcp subservice the struct/api is residing (gcp is split into service.subservice.list)
	SubService string
	// ListFunction function string which lists all resources
	ListFunction string
	// OutputField is field where the result is located in the output struct
	OutputField string
	// Template is the template to use to generate the resource (some services has different template as some services were generated using different original codegen)
	Template string
	// imports to add for this resource
	Imports []string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	// Columns override, override generated columns
	OverrideColumns []codegen.ColumnDefinition
}

var ProjectIdColumn = codegen.ColumnDefinition{
	Name:     "project_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveProject",
}
