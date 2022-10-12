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
	// CFStruct that will be used to generate the cloudquery table
	CFStruct interface{}
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	PrimaryKey string
	Template   string
	Multiplex  string

	DefaultColumns []codegen.ColumnDefinition
	ExtraColumns   []codegen.ColumnDefinition
	RenameColumns  map[string]string

	Relations        []string
	ResolverFuncName string
	TableFuncName    string
	Filename         string
	Package          string
	ImportClient     bool
}

var (
	AccountIDColumn = codegen.ColumnDefinition{
		Name:        "account_id",
		Description: "The Account ID of the resource.",
		Type:        schema.TypeString,
		Resolver:    "client.ResolveAccountID",
	}

	ZoneIDColumn = codegen.ColumnDefinition{
		Name:        "zone_id",
		Description: "Zone identifier tag.",
		Type:        schema.TypeString,
		Resolver:    "client.ResolveZoneID",
	}
)
