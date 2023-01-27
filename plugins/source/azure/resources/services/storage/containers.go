package storage

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func containers() *schema.Table {
	return &schema.Table{
		Name:        "azure_storage_containers",
		Resolver:    fetchContainers,
		Description: "https://learn.microsoft.com/en-us/rest/api/storagerp/blob-containers/list?tabs=HTTP#listcontaineritem",
		Transform:   transformers.TransformWithStruct(&armstorage.ListContainerItem{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{},
	}
}
