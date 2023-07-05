package billing_history

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func BillingHistory() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_billing_history",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/billingHistory_list",
		Resolver:    fetchBillingHistoryBillingHistory,
		Transform:   transformers.TransformWithStruct(&godo.BillingHistoryEntry{}),
		Columns:     []schema.Column{},
	}
}
