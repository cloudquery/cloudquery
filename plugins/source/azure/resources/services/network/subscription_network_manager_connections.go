package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SubscriptionNetworkManagerConnections() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_subscription_network_manager_connections",
		Resolver:    fetchSubscriptionNetworkManagerConnections,
		Description: "https://learn.microsoft.com/en-us/rest/api/networkmanager/management-group-network-manager-connections/list?tabs=HTTP#networkmanagerconnection",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_subscription_network_manager_connections", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.ManagerConnection{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
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

func fetchSubscriptionNetworkManagerConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewSubscriptionNetworkManagerConnectionsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
