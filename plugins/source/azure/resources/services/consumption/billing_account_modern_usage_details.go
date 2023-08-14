package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func BillingAccountModernUsageDetails() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_billing_account_modern_usage_details",
		Resolver:             fetchBillingAccountModernUsageDetails,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/usage-details/list?tabs=HTTP#modernusagedetail",
		Multiplex:            client.ModernBillingAccountMultiplex,
		Transform:            transformers.TransformWithStruct(&armconsumption.ModernUsageDetail{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchBillingAccountModernUsageDetails(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewUsageDetailsClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(*cl.BillingAccount.ID, nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		for _, v := range p.Value {
			res <- v.(*armconsumption.ModernUsageDetail)
		}
	}
	return nil
}
