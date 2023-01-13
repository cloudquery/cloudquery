package billing_history

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func BillingHistory() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_billing_history",
		Resolver:  fetchBillingHistoryBillingHistory,
		Transform: transformers.TransformWithStruct(&godo.BillingHistoryEntry{}),
		Columns:   []schema.Column{},
	}
}
