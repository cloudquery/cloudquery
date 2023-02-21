package desktopvirtualization

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func HostPools() *schema.Table {
	return &schema.Table{
		Name:        "azure_desktopvirtualization_host_pools",
		Resolver:    fetchHostPools,
		Description: "https://learn.microsoft.com/en-us/rest/api/desktopvirtualization/host-pools/list?tabs=HTTP#hostpool",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_desktopvirtualization_host_pools", client.Namespacemicrosoft_desktopvirtualization),
		Transform:   transformers.TransformWithStruct(&armdesktopvirtualization.HostPool{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.SubscriptionID},
	}
}

func fetchHostPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdesktopvirtualization.NewHostPoolsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
