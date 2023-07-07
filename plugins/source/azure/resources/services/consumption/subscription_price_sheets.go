package consumption

import (
	"context"
	"net/url"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func SubscriptionPriceSheets() *schema.Table {
	return &schema.Table{
		Name:                 "azure_consumption_subscription_price_sheets",
		Resolver:             fetchSubscriptionPriceSheets,
		PostResourceResolver: client.LowercaseIDResolver,
		Description:          "https://learn.microsoft.com/en-us/rest/api/consumption/price-sheet/get?tabs=HTTP#pricesheetresult",
		Multiplex:            client.SubscriptionMultiplexRegisteredNamespace("azure_consumption_subscription_price_sheets", client.Namespacemicrosoft_consumption),
		Transform:            transformers.TransformWithStruct(&armconsumption.PriceSheetResult{}, transformers.WithPrimaryKeys("ID")),
	}
}

func fetchSubscriptionPriceSheets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armconsumption.NewPriceSheetClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	opts := &armconsumption.PriceSheetClientGetOptions{Expand: to.Ptr("properties/meterDetails")}
	resp, err := svc.Get(ctx, opts)
	if err != nil {
		return err
	}

	// This is a workaround to get all price sheets
	// Somehow related to https://github.com/Azure/azure-sdk-for-go/issues/4142
	allPricesSheets := resp.PriceSheetResult.Properties.Pricesheets
	nextLink := resp.PriceSheetResult.Properties.NextLink
	for nextLink != nil {
		parsedNextLink, err := url.Parse(*nextLink)
		if err != nil {
			cl.Logger().Warn().Err(err).Msgf("failed to parse next link: %q", *nextLink)
			break
		}
		token := parsedNextLink.Query().Get("skiptoken")
		if token == "" {
			cl.Logger().Warn().Msgf("failed to get skiptoken from next link: %q", *nextLink)
			break
		}
		opts.Skiptoken = to.Ptr(token)
		paginatedResponse, err := svc.Get(ctx, opts)
		if err != nil {
			cl.Logger().Warn().Err(err).Msgf("failed to get paginated response for next link: %q", *nextLink)
			break
		}
		allPricesSheets = append(allPricesSheets, paginatedResponse.PriceSheetResult.Properties.Pricesheets...)
		nextLink = paginatedResponse.PriceSheetResult.Properties.NextLink
	}

	resp.PriceSheetResult.Properties.Pricesheets = allPricesSheets
	res <- resp.PriceSheetResult
	return nil
}
