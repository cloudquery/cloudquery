package balances

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Balances() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_balances",
		Resolver:  fetchBalancesBalances,
		Transform: transformers.TransformWithStruct(&godo.Balance{}),
		Columns:   []schema.Column{},
	}
}
