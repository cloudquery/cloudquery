package codegen

import (
	"fmt"
	"log"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
)

type Resource struct {
	// DefaultColumns columns that will be appended to the main table
	DefaultColumns []codegen.ColumnDefinition
	// Table is the table definition that will be used to generate the CloudQuery table
	Table *codegen.TableDefinition
	// AzureStruct that will be used to generate the CloudQuery table
	AzureStruct interface{}
	// AzureStructName is the name of the AzureStruct because it can't be inferred by reflection
	AzureStructName string
	// AzureService is the name of the azure service the struct/api is residing
	AzureService string
	// AzureSubService is the name of the azure subservice the struct/api is residing
	AzureSubService string
	// Template is the templates to use to generate the resource
	Templates []string
	// imports to add for this resource
	Imports []string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	//CreateTableOptions options to use to create the main table
	CreateTableOptions schema.TableCreationOptions
	// List or ListAll
	ListFunction string
}

const pluginName = "azure"

var SubscriptionIdColumn = codegen.ColumnDefinition{
	Name:     "subscription_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveAzureSubscription",
}

func initResourceTable(resource *Resource) {
	var err error
	resource.Table, err = codegen.NewTableFromStruct(fmt.Sprintf("%s_%s_%s", pluginName, resource.AzureService, resource.AzureSubService), resource.AzureStruct, codegen.WithSkipFields(resource.SkipFields...))
	if err != nil {
		log.Fatal(err)
	}
	resource.Table.Columns = append(resource.DefaultColumns, resource.Table.Columns...)
	resource.Table.Multiplex = "client.SubscriptionMultiplex"
	resource.Table.Resolver = "fetch" + strcase.ToCamel(resource.AzureService) + strcase.ToCamel(resource.AzureSubService)
	resource.Table.Options.PrimaryKeys = resource.CreateTableOptions.PrimaryKeys
}
