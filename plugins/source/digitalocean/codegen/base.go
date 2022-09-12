package codegen

import (
	"github.com/cloudquery/plugin-sdk/codegen"
)

type Resource struct {
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// Struct that will be used to generate the cloudquery table
	Struct interface{}
	// MockStruct that will be used to generate the mock
	IsStructPointer bool
	MockStruct      interface{}
	// ParentStruct that will be used to generate the mock
	ParentStruct interface{}
	// ParentStructName that will be used to generate the mock
	IsParentPointer bool
	// Args for get list function
	Args             string
	ParentStructName string
	// MockWrapper
	MockWrapper bool
	// MockStructName is the name of the Struct because it can't be inferred by reflection
	MockStructName string
	// StructName is the name of the Struct because it can't be inferred by reflection
	StructName string
	// Service is the name of the gcp service the struct/api is residing
	Service string
	// SubService is the name of the subservice
	SubService string
	// SubService is the name of the subservice
	SubServiceName string
	Relations      []string
	// Template is the template to use to generate the resource (some services has different template as some services were generated using different original codegen)
	Template string
	// imports to add for this resource
	Imports []string
	// Multiplex
	Multiplex *string
	// ChildTable
	ChildTable bool
	// SkipMock is used to skip the mock generation for this resource
	SkipMock bool
	// Pass to MockTemplate
	MockTemplate string
	// MockFieldName is a name of a struct that is used for mocking
	ResponsePath string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	// Columns override, override generated columns
	ExtraColumns []codegen.ColumnDefinition
	FunctionName string
	SkipFetch    bool
}
