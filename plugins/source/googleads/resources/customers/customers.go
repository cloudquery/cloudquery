package customers

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/cloudquery/plugins/source/googleads/gaql"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/shenzhencenter/google-ads-pb/enums"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

var customerOptions = &gaql.Options{
	Expand: []string{"CallReportingSetting", "ConversionTrackingSetting", "RemarketingSetting"},
}

func Customers() *schema.Table {
	return &schema.Table{
		Name:        "googleads_customers",
		Description: client.APIRef + "/Customer",
		Multiplex:   client.MultiplexByCustomer,
		Transform: client.TransformWithStruct(new(resources.Customer),
			transformers.WithPrimaryKeys("Id", "ResourceName"),
		),
		Resolver:  fetchCustomers,
		Relations: schema.Tables{customerLabels()},
	}
}

func fetchCustomers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	ctx = c.OutgoingContext(ctx)

	req := &services.SearchGoogleAdsStreamRequest{
		CustomerId: c.CustomerID,
		Query: gaql.Query(new(resources.Customer), nil, customerOptions) +
			"\nWHERE customer.id = " + c.CustomerID,
		SummaryRowSetting: enums.SummaryRowSettingEnum_NO_SUMMARY_ROW,
	}

	resp, err := c.GoogleAdsClient.SearchStream(ctx, req)
	if err != nil {
		return err
	}

	return client.ReceiveStream(resp.Recv, (*services.GoogleAdsRow).GetCustomer, res)
}
