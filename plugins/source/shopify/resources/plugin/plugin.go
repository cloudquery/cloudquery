package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/services/checkout"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/services/customer"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/services/order"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/services/price_rule"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/resources/services/product"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"shopify",
		Version,
		[]*schema.Table{
			customer.Customers(),
			order.Orders(),
			product.Products(),
			price_rule.PriceRules(),
			checkout.AbandonedCheckouts(),
		},
		client.Configure,
	)
}
