package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func PublicIpAddresses() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_public_ip_addresses",
		Resolver:    fetchPublicIpAddresses,
		Description: "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/public-ip-addresses/list?tabs=HTTP#publicipaddress",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_public_ip_addresses", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.PublicIPAddress{}),
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

func fetchPublicIpAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewPublicIPAddressesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
