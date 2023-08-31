package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func VpnServerConfigurations() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_vpn_server_configurations",
		Resolver:             fetchVpnServerConfigurations,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/virtualwan/vpn-server-configurations/list?tabs=HTTP#vpnserverconfiguration",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_vpn_server_configurations", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.VPNServerConfiguration{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchVpnServerConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewVPNServerConfigurationsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
