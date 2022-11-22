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

	Relations        []string
	ResolverFuncName string
	TableFuncName    string
	Filename         string
	Package          string
	ImportClient     bool
}

var (
	SharingIDColumn = codegen.ColumnDefinition{
		Name:        "sharing_id",
		Description: "The Sharing ID of the resource.",
		Type:        schema.TypeString,
		Resolver:    "client.ResolveSharingID",
	}
)
