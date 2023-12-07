package storage

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func tables() *schema.Table {
	return &schema.Table{
		Name:                 "azure_storage_tables",
		Resolver:             fetchTables,
		PostResourceResolver: client.LowercaseIDResolver,
		PreResourceResolver:  getTable,
		Description:          "https://learn.microsoft.com/en-us/rest/api/storagerp/table/list?tabs=HTTP#table",
		Transform:            transformers.TransformWithStruct(&armstorage.Table{}, transformers.WithPrimaryKeys("ID"), transformers.WithSkipFields("TableProperties")),
		Columns: schema.ColumnList{
			client.SubscriptionID,
			{
				Name:     "properties",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("TableProperties"),
			},
		},
	}
}
