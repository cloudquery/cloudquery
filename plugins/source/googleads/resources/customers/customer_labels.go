package customers

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleads/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/shenzhencenter/google-ads-pb/resources"
	"github.com/shenzhencenter/google-ads-pb/services"
)

func customerLabels() *schema.Table {
	return &schema.Table{
		Name:        "googleads_customer_labels",
		Description: client.APIRef + "/CustomerLabel",
		Transform: client.TransformWithStruct(new(resources.CustomerLabel),
			transformers.WithPrimaryKeys("Customer", "ResourceName"),
		),
		Columns:  schema.ColumnList{client.CustomerID},
		Resolver: client.Fetcher((*services.GoogleAdsRow).GetCustomerLabel),
	}
}
