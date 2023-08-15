package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Watchers() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_watchers",
		Resolver:             fetchWatchers,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/network-watcher/network-watchers/list-all?tabs=HTTP#networkwatcher",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_watchers", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.Watcher{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
		Relations: []*schema.Table{
			watcherFlowLogs(),
		},
	}
}

func fetchWatchers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewWatchersClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
