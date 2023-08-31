package storage

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func containers() *schema.Table {
	return &schema.Table{
		Name:                 "azure_storage_containers",
		Resolver:             fetchContainers,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/storagerp/blob-containers/list?tabs=HTTP#listcontaineritem",
		Transform:            transformers.TransformWithStruct(&armstorage.ListContainerItem{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}
