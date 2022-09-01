package recipes

import (
	"fmt"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

const ResolverAuto = "auto"

type Resource struct {
	// DefaultColumns columns that will be appended to the main table
	DefaultColumns []codegen.ColumnDefinition
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// AWSStruct that will be used to generate the cloudquery table
	AWSStruct interface{}

	// AWSService is the name of the aws service the struct/api is residing. Capitalization is important as it's also used in the client's service map.
	AWSService string

	// Template is the template to use to generate the resource
	Template string

	MultiplexerServiceOverride string
	CQSubserviceOverride       string // used in table and file names

	PaginatorStruct    interface{} // Used only in resource_list_and_detail and list_describe templates.
	PaginatorGetStruct interface{} // Used only in resource_list_and_detail and list_describe templates.

	ItemName                string      // Override. Defaults to AWSStructName
	AWSSubService           string      // Override. Name of the aws subservice the struct/api is residing. Should be in CamelCase. Inferred from ItemsStruct.
	ItemsStruct             interface{} // This should point to a verb + .AWSSubService + "Output"
	ItemsCustomOptionsBlock string      // Only supported by resource_get template for now.

	CustomErrorBlock string // Only used in list_and_detail template.
	CustomTagField   string // Only used in list_and_detail template.

	Parent          *Resource
	ParentFieldName string
	ChildFieldName  string // Override. Defaults to ParentFieldName

	SkipDescribeParentInputs bool // Only used in list_describe template.

	// imports to add for this resource
	Imports     []string
	MockImports []string

	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	//CreateTableOptions options to use to create the main table
	CreateTableOptions schema.TableCreationOptions

	ColumnOverrides map[string]codegen.ColumnDefinition

	AddTypesImport bool   // add types import regardless of template spec (can lead to double imports)
	TrimPrefix     string // trim this prefix from all column names

	CustomInputs  []string // Custom inputs to the first call (Get or List)
	CustomInputs2 []string // Custom inputs to the second call (Describe after a List)

	AutoCalculated
}

type AutoCalculated struct {
	AWSStructName     string // Automatically resolved using reflection from AWSStruct
	PaginatorListName string // Auto calculated from PaginatorStruct.
	PaginatorListType string // Auto calculated from PaginatorStruct.
	ResponseItemsName string // Auto calculated from ItemsStruct by default, otherwise defaults to Items
	NextTokenName     string // Auto calculated from ItemsStruct for resource_get template.

	ListMethod string // Auto calculated from PaginatorStruct
	GetMethod  string // Auto calculated from ItemsStruct

	// Field matchers
	MatchedGetAndListFields map[string]string // field name -> "item." + field name (or sometimes just "&item")
	GetAndListOrder         []string          // Order of fields for consistency

	HasTags bool // autodetected by scanning all columns for `tags`

	TableFuncName string
	MockFuncName  string
	TestFuncName  string
	NestingLevel  int
	TypesImport   string

	TemplateFilename string // This is injected to top of every template result
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

	AllResources []*Resource
)

func add(list ...*Resource) {
	AllResources = append(AllResources, list...)
}

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
