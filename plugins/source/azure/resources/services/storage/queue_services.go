package storage

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func queueServices() *schema.Table {
	return &schema.Table{
		Name:                 "azure_storage_queue_services",
		Resolver:             fetchQueueServices,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/storagerp/queue-services/list?tabs=HTTP#queueserviceproperties",
		Transform:            transformers.TransformWithStruct(&armstorage.QueueServiceProperties{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchQueueServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armstorage.NewQueueServicesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	item := parent.Item.(*armstorage.Account)
	group, err := client.ParseResourceGroup(*item.ID)
	if err != nil {
		return err
	}
	resp, err := svc.List(ctx, group, *item.Name, nil)
	if err != nil {
		return err
	}
	res <- resp.Value
	return nil
}
