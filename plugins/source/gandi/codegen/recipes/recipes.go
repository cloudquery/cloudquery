package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

type Resource struct {
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// DataStruct that will be used to generate the cloudquery table
	DataStruct interface{}
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	Template   string
	Multiplex  string

	ExtraColumns []codegen.ColumnDefinition
	PKColumns    []string

	PreResourceResolver   string
	PostResourceResolver  string
	Relations             []string
	UnwrapEmbeddedStructs bool

	// These are inferred with reflection but can be overridden
	Service    string // Inferred from DataStruct package, pluralized
	SubService string // Inferred from DataStruct name, singular
	TableName  string // singular Service + plural SubService

	// These are auto calculated
	ImportClient     bool   // true if the resource/column resolvers use the client package
	Filename         string // Calculated from TableName
	TableFuncName    string // Calculated from TableName
	ResolverFuncName string // Calculated from TableFuncName

	//used for generating better table names
	//parent   *Resource
	//children []*Resource
}

var (
	SharingIDColumn = codegen.ColumnDefinition{
		Name:        "sharing_id",
		Description: "The Sharing ID of the resource.",
		Type:        schema.TypeString,
		Resolver:    "client.ResolveSharingID",
	}
)
