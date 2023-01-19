package storage

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func blob_services() *schema.Table {
	return &schema.Table{
		Name:        "azure_storage_blob_services",
		Resolver:    fetchBlobServices,
		Description: "https://learn.microsoft.com/en-us/rest/api/storagerp/blob-services/list?tabs=HTTP#blobserviceproperties",
		Transform:   transformers.TransformWithStruct(&armstorage.BlobServiceProperties{}),
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
