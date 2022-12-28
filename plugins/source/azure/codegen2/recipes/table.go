package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

var SubscriptionIdColumn = codegen.ColumnDefinition{
	Name:     "subscription_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveAzureSubscription",
}

var DefaultExtraColumns = []codegen.ColumnDefinition{
	SubscriptionIdColumn,
}

var Tables []Table

type Table struct {
	// PackageName name is the packgename in the source plugin this resource is located
	PackageName string
	// Sets PreResourceResolver
	PreResourceResolver string
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// Struct that will be used to generate the cloudquery table
	Struct any
	// StructName is the name of the Struct because it can't be inferred by reflection
	StructName string
	// ImportPath need for each table
	ImportPath string
	// Service is the name of the azure service the struct/api is residing
	Service string
	// Name is the name of the azure subservice the struct/api is residing (gcp is split into service.subservice.list)
	Name string
	// NewFunction
	Client any
	// ClientNameName name of the above function via Reflection
	ClientName string
	// ResponseStruct
	ResponseStruct any
	// ResponseStructName is reflected name from the ResponseStruct
	ResponseStructName string
	// Does responseStruct has NextLink field
	ResponspeStructNextLink bool
	// NewFunc
	NewFunc any
	// NewFuncName is reflected name from the ResponseStruct
	NewFuncName string
	// Does the newfunc get subscription_id as a first parameter
	NewFuncHasSubscriptionId bool
	// ListFunc
	ListFunc any
	// NewFuncName is reflected name from the ListFunc
	ListFuncName string
	// Does the ListFunc get subscription_id as a first parameter
	// ListFuncHasResourceGroupName bool
	// Relations is list of relations functions
	Relations []*Table
	// Multiplex
	Multiplex string
	// ChildTable
	ChildTable bool
	// SkipMock is used to skip the mock generation for this resource
	SkipMock bool
	// Don't generate fetch
	SkipFetch bool
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	// URL is the rest endpoint. This is used by mock tests
	URL          string
	ExtraColumns []codegen.ColumnDefinition
}
