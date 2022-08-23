package codegen

import (
	"fmt"
	"log"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
)

type Template struct {
	Source      string
	Destination string
}

type Resource struct {
	// DefaultColumns columns that will be appended to the main table
	DefaultColumns []codegen.ColumnDefinition
	// Table is the table definition that will be used to generate the CloudQuery table
	Table *codegen.TableDefinition
	// AzureStruct that will be used to generate the CloudQuery table
	AzureStruct        interface{}
	AzureStructName    string
	AzurePackageName   string
	AzureService       string
	AzureSubService    string
	Templates          []Template
	Imports            []string
	MockImports        []string
	SkipFields         []string
	CreateTableOptions schema.TableCreationOptions
	// Used in the `resource_list` templates
	ListFunction string
}

const pluginName = "azure"

var SubscriptionIdColumn = codegen.ColumnDefinition{
	Name:     "subscription_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveAzureSubscription",
}

var IdColumn = codegen.ColumnDefinition{
	Name:     "id",
	Type:     schema.TypeString,
	Resolver: "schema.PathResolver(\"ID\")",
}

func initResourceTable(resource *Resource) {
	var err error
	resource.Table, err = codegen.NewTableFromStruct(fmt.Sprintf("%s_%s_%s", pluginName, resource.AzurePackageName, strcase.ToSnake(resource.AzureSubService)), resource.AzureStruct, codegen.WithSkipFields(resource.SkipFields...))
	if err != nil {
		log.Fatal(err)
	}
	resource.Table.Columns = append(resource.DefaultColumns, resource.Table.Columns...)
	resource.Table.Multiplex = "client.SubscriptionMultiplex"
	resource.Table.Resolver = "fetch" + resource.AzureService + resource.AzureSubService
	resource.Table.Options.PrimaryKeys = resource.CreateTableOptions.PrimaryKeys
}
