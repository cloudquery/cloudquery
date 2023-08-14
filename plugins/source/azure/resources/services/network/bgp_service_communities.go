package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BgpServiceCommunities() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_bgp_service_communities",
		Resolver:             fetchBgpServiceCommunities,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/expressroute/bgp-service-communities/list?tabs=HTTP#bgpservicecommunity",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_bgp_service_communities", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.BgpServiceCommunity{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionIDPK},
	}
}

func fetchBgpServiceCommunities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewBgpServiceCommunitiesClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
