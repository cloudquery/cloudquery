package storage

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func tables() *schema.Table {
	return &schema.Table{
		Name:        "azure_storage_tables",
		Resolver:    fetchTables,
		Description: "https://learn.microsoft.com/en-us/rest/api/storagerp/table/list?tabs=HTTP#table",
		Transform:   transformers.TransformWithStruct(&armstorage.Table{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
