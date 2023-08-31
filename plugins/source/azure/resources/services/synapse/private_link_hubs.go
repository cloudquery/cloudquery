package synapse

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func PrivateLinkHubs() *schema.Table {
	return &schema.Table{
		Name:                 "azure_synapse_private_link_hubs",
		Resolver:             fetchPrivateLinkHubs,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/synapse/private-link-hubs/list?tabs=HTTP#privatelinkhub",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_synapse_private_link_hubs", client.Namespacemicrosoft_synapse),
		Transform:            transformers.TransformWithStruct(&armsynapse.PrivateLinkHub{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
	}
}

func fetchPrivateLinkHubs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armsynapse.NewPrivateLinkHubsClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
