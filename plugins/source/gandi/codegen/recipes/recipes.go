package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

type Resource struct {
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// TableName can be used to override the default generated table name
	TableName string
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

	// These are inferred with reflection
	Package string // Inferred from DataStruct package name pluralized

	// These are auto calculated
	ImportClient     bool   // true if the resource/column resolvers use the client package
	Filename         string // Calculated from TableName
	TableFuncName    string // Calculated from TableName
	ResolverFuncName string // Calculated from TableFuncName

	// used for generating resolver and mock tests, but set automatically
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
