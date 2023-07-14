package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PrivateLinkServices() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_private_link_services",
		Resolver:             fetchPrivateLinkServices,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/private-link-services/list-by-subscription?tabs=HTTP#privatelinkservice",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_private_link_services", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.PrivateLinkService{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPrivateLinkServices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewPrivateLinkServicesClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListBySubscriptionPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
