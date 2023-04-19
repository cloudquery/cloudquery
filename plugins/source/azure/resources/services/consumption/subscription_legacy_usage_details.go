package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func SubscriptionLegacyUsageDetails() *schema.Table {
	return &schema.Table{
		Name:        "azure_consumption_subscription_legacy_usage_details",
		Resolver:    fetchSubscriptionLegacyUsageDetails,
		Description: "https://learn.microsoft.com/en-us/rest/api/consumption/usage-details/list?tabs=HTTP#legacyusagedetail",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_consumption_subscription_legacy_usage_details", client.Namespacemicrosoft_consumption),
		Transform:   transformers.TransformWithStruct(&armconsumption.LegacyUsageDetail{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchSubscriptionLegacyUsageDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewUsageDetailsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	scope := "subscriptions/" + cl.SubscriptionId
	pager := svc.NewListPager(scope, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, v := range p.Value {
			res <- v.(*armconsumption.LegacyUsageDetail)
		}
	}
	return nil
}
