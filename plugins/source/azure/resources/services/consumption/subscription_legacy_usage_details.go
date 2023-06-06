package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/cloudquery/plugin-sdk/v3/types"
)

type usageDetail struct {
	*armconsumption.UsageDetail

	Properties any `json:"properties"`
}

func SubscriptionLegacyUsageDetails() *schema.Table {
	return &schema.Table{
		Name:        "azure_consumption_subscription_legacy_usage_details",
		Resolver:    fetchSubscriptionLegacyUsageDetails,
		Description: "https://learn.microsoft.com/en-us/rest/api/consumption/usage-details/list?tabs=HTTP#legacyusagedetail",
		Multiplex:   client.SubscriptionBillingPeriodMultiplex,
		Transform: transformers.TransformWithStruct(
			&usageDetail{},
			transformers.WithPrimaryKeys("ID"),
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Columns: schema.ColumnList{
			{
				Name:     "properties",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("Properties"),
			},
		},
	}
}

func fetchSubscriptionLegacyUsageDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewUsageDetailsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(*cl.BillingPeriod.ID, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, v := range p.Value {
			ud := &usageDetail{UsageDetail: v.GetUsageDetail()}
			switch d := v.(type) {
			case *armconsumption.LegacyUsageDetail:
				ud.Properties = d.Properties
			case *armconsumption.ModernUsageDetail:
				ud.Properties = d.Properties
			}
			res <- ud
		}
	}
	return nil
}
