package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func RestorePointCollections() *schema.Table {
	return &schema.Table{
		Name:                 "azure_compute_restore_point_collections",
		Resolver:             fetchRestorePointCollections,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/compute/restore-point-collections/list?tabs=HTTP#restorepointcollection",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_compute_restore_point_collections", client.Namespacemicrosoft_compute),
		Transform:            transformers.TransformWithStruct(&armcompute.RestorePointCollection{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchRestorePointCollections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armcompute.NewRestorePointCollectionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListAllPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
