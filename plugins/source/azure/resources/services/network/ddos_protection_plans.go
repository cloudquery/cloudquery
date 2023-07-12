package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func DdosProtectionPlans() *schema.Table {
	return &schema.Table{
		Name:                 "azure_network_ddos_protection_plans",
		Resolver:             fetchDdosProtectionPlans,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/virtualnetwork/ddos-protection-plans/list?tabs=HTTP#ddosprotectionplan",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_network_ddos_protection_plans", client.Namespacemicrosoft_network),
		Transform:            transformers.TransformWithStruct(&armnetwork.DdosProtectionPlan{}, transformers.WithPrimaryKeys("ID")),
		Columns:              schema.ColumnList{client.SubscriptionID},
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
