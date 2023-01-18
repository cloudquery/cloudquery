package networkfunction

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func AzureTrafficCollectorsBySubscription() *schema.Table {
	return &schema.Table{
		Name:      "azure_networkfunction_azure_traffic_collectors_by_subscription",
		Resolver:  fetchAzureTrafficCollectorsBySubscription,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_networkfunction_azure_traffic_collectors_by_subscription", client.Namespacemicrosoft_networkfunction),
		Transform: transformers.TransformWithStruct(&armnetworkfunction.AzureTrafficCollector{}),
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

func fetchAzureTrafficCollectorsBySubscription(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetworkfunction.NewAzureTrafficCollectorsBySubscriptionClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
