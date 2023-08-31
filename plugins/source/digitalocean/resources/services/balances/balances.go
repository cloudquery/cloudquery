package balances

import (
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func Balances() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_balances",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/balance_get",
		Resolver:    fetchBalancesBalances,
		Transform:   transformers.TransformWithStruct(&godo.Balance{}),
		Columns:     []schema.Column{},
	}
}
