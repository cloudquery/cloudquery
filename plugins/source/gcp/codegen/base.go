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
	// GCPStruct that will be used to generate the cloudquery table
	GCPStruct interface{}
	// GCPStructName is the name of the GCPStruct because it can't be inferred by reflection
	GCPStructName string
	// GCPService is the name of the gcp service the struct/api is residing
	GCPService string
	// GCPSubService is the name of the gcp subservice the struct/api is residing (gcp is split into service.subservice.list)
	GCPSubService string
	// Template is the template to use to generate the resource (some services has different template as some services were generated using different original codegen)
	Template string
	// imports to add for this resource
	Imports []string
	// MockImports imports to add for mock tests
	MockImports []string
	// MockListStruct specified the name of the returned list function. There are
	// some inconsistencies in naming so we have to have a way of manually overriding defaults
	MockListStruct string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	//CreateTableOptions options to use to create the main table
	CreateTableOptions schema.TableCreationOptions
}

var ProjectIdColumn = codegen.ColumnDefinition{
	Name:     "project_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveProject",
}
