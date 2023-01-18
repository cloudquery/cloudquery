package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DdosProtectionPlans() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_ddos_protection_plans",
		Resolver:    fetchDdosProtectionPlans,
		Description: "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/ddos-protection-plans/list?tabs=HTTP#ddosprotectionplan",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_network_ddos_protection_plans", client.Namespacemicrosoft_network),
		Transform:   transformers.TransformWithStruct(&armnetwork.DdosProtectionPlan{}),
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

func fetchDdosProtectionPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armnetwork.NewDdosProtectionPlansClient(cl.SubscriptionId, cl.Creds, cl.Options)
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
