package recipes

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

type Resource struct {
	// PackageName name is the packagename in the source plugin this resource is located
	//PackageName string
	// DefaultColumns columns that will be appended to the main table
	DefaultColumns []codegen.ColumnDefinition
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// AWSStruct that will be used to generate the cloudquery table
	AWSStruct interface{}
	// AWSStructName is the name of the AWSStruct, if necessary (automatically resolved using reflection from AWSStruct)
	AWSStructName string
	// AWSService is the name of the aws service the struct/api is residing. Capitalization is important as it's also used in the client's service map.
	AWSService string
	// AWSSubService is the name of the aws subservice the struct/api is residing. Should be in CamelCase
	AWSSubService string
	// Template is the template to use to generate the resource (some services has different template as some services were generated using different original codegen)
	Template string

	MultiplexerServiceOverride string
	CQSubserviceOverride       string // used in table and file names

	ListVerb      string // Override. Defaults to "List". Only used in list_describe template.
	ListFieldName string // Only used in list_describe template.

	ItemName          string // Override. Defaults to AWSStructName
	Verb              string // Override. Default depends on template used.
	ResponseItemsName string // Override. Defaults to Items

	Parent          *Resource
	ParentFieldName string

	// imports to add for this resource
	Imports []string
	// MockImports imports to add for mock tests
	MockImports []string
	// MockListStruct specified the name of the returned list function. There are
	// some inconsistencies in naming, so we have to have a way of manually overriding defaults
	MockListStruct string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	//CreateTableOptions options to use to create the main table
	CreateTableOptions schema.TableCreationOptions

	ColumnOverrides map[string]codegen.ColumnDefinition

	HasTags         bool // autodetected by scanning all columns for `tags`
	SkipTypesImport bool // skip "types" import (except for mock mode)
	AddTypesImport  bool // always add "types" import

	TrimPrefix string // trim this prefix from all column names

	TableFuncName string // auto calculated
	MockFuncName  string // auto calculated
	TestFuncName  string // auto calculated
	NestingLevel  int    // auto calculated

	CustomResolvers []string
	CustomInputs    []string
}

var (
	AccountIdColumn = codegen.ColumnDefinition{
		Name:        "account_id",
		Description: "The AWS Account ID of the resource.",
		Type:        schema.TypeString,
		Resolver:    "client.ResolveAWSAccount",
	}
	RegionColumn = codegen.ColumnDefinition{
		Name:        "region",
		Description: "The AWS Region of the resource.",
		Type:        schema.TypeString,
		Resolver:    "client.ResolveAWSRegion",
	}
	NamespaceColumn = codegen.ColumnDefinition{
		Name:        "namespace",
		Description: "The AWS Service Namespace of the resource.",
		Type:        schema.TypeString,
		Resolver:    "client.ResolveAWSNamespace",
	}
)

// parentize adds the given parent to each resource (in subs) and returns the combined list
func parentize(parent *Resource, subs ...*Resource) []*Resource {
	ret := make([]*Resource, len(subs)+1)
	ret[0] = parent
	for i := range subs {
		if subs[i].Parent == nil {
			subs[i].Parent = parent
		}
		subs[i].NestingLevel++
		if subs[i].AWSService == "" {
			subs[i].AWSService = subs[i].Parent.AWSService
		}
		ret[i+1] = subs[i]
	}
	return ret
}

// combine the given *Resource or []*Resource into a single []*Resource
// if the given argument is of another type, combine will panic
func combine(list ...interface{}) []*Resource {
	res := make([]*Resource, 0, len(list))
	for i := range list {
		switch v := list[i].(type) {
		case *Resource:
			res = append(res, v)
		case []*Resource:
			res = append(res, v...)
		default:
			panic(fmt.Sprintf("combine: unhandled type %T", list[i]))
		}
	}
	return res
}
