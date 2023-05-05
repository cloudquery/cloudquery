package billing_history

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func BillingHistory() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_billing_history",
		Description: "https://pkg.go.dev/github.com/digitalocean/godo#BillingHistoryEntry",
		Resolver:    fetchBillingHistoryBillingHistory,
		Transform:   transformers.TransformWithStruct(&godo.BillingHistoryEntry{}),
		Columns:     []schema.Column{},
	}
}
