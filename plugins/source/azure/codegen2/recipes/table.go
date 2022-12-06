package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
)

var Tables []Table

type Table struct {
	// PackageName name is the packgename in the source plugin this resource is located
	PackageName string
	// Sets PreResourceResolver
	PreResourceResolver string
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// Struct that will be used to generate the cloudquery table
	Struct interface{}
	// StructName is the name of the Struct because it can't be inferred by reflection
	StructName string
	// ImportPath need for each table
	ImportPath string
	// Service is the name of the azure service the struct/api is residing
	Service string
	// Name is the name of the azure subservice the struct/api is residing (gcp is split into service.subservice.list)
	Name string
	// NewFunction
	Client interface{}
	// ClientNameName name of the above function via Reflection
	ClientName string
	// ResponseStruct
	ResponseStruct interface{}
	// ResponseStructName is reflected name from the ResponseStruct
	ResponseStructName string
	// Does responseStruct has NextLink field
	ResponspeStructNextLink bool
	// NewFunc
	NewFunc interface{}
	// NewFuncName is reflected name from the ResponseStruct
	NewFuncName string
	// Does the newfunc get subscription_id as a first parameter
	NewFuncHasSubscriptionId bool
	// ListFunc
	ListFunc interface{}
	// NewFuncName is reflected name from the ListFunc
	ListFuncName string
	// Does the ListFunc get subscription_id as a first parameter
	ListFuncHasSubscriptionId bool
	// should we use ListAll function
	ListAll bool
	// Relations is list of relations functions
	Relations []string
	// Template is the template to use to generate the resource (some services has different template as some services were generated using different original codegen)
	Template string
	// imports to add for this resource
	// Imports []string
	// Multiplex
	Multiplex *string
	// ChildTable
	ChildTable bool
	// FakerFieldsToIgnore is a list of fields to ignore when generating faker data
	FakerFieldsToIgnore []string
	// SkipMock is used to skip the mock generation for this resource
	SkipMock bool
	// Pass to MockTemplate
	MockTemplate string
	// MockListStruct is the name of the struct that will be used in the mock template
	MockListStruct string
	// MockImports imports used in mock tests
	// MockImports []string
	// Don't generate fetch
	SkipFetch bool
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	// ExtraColumns override, override generated columns
	ExtraColumns []codegen.ColumnDefinition
	URL          string
}
