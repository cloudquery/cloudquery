package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/customers"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/disputes"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/invoices"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/products"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/refunds"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/resources/services/subscriptions"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"stripe",
		Version,
		[]*schema.Table{
			customers.Customers(),
			disputes.Disputes(),
			invoices.Invoices(),
			invoices.InvoiceItems(),
			products.Products(),
			refunds.Refunds(),
			subscriptions.Subscriptions(),
		},
		client.Configure,
	)
}
