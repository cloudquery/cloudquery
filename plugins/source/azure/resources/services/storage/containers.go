package storage

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func containers() *schema.Table {
	return &schema.Table{
		Name:      "azure_storage_containers",
		Resolver:  fetchContainers,
		Transform: transformers.TransformWithStruct(&armstorage.ListContainerItem{}),
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
