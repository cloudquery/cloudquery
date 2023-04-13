package consumption

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SubscriptionPriceSheets() *schema.Table {
	return &schema.Table{
		Name:        "azure_consumption_subscription_price_sheets",
		Resolver:    fetchSubscriptionPriceSheets,
		Description: "https://learn.microsoft.com/en-us/rest/api/consumption/price-sheet/get?tabs=HTTP#pricesheetresult",
		Multiplex:   client.SubscriptionMultiplexRegisteredNamespace("azure_consumption_subscription_price_sheets", client.Namespacemicrosoft_consumption),
		Transform:   transformers.TransformWithStruct(&armconsumption.PriceSheetResult{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchSubscriptionPriceSheets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewPriceSheetClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	resp, err := svc.Get(ctx, &armconsumption.PriceSheetClientGetOptions{Expand: to.Ptr("properties/meterDetails")})
	if err != nil {
		return err
	}
	res <- resp.PriceSheetResult
	return nil
}
