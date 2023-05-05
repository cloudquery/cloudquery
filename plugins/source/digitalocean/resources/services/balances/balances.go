package balances

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func Balances() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_balances",
		Description: "https://pkg.go.dev/github.com/digitalocean/godo#Balance",
		Resolver:    fetchBalancesBalances,
		Transform:   transformers.TransformWithStruct(&godo.Balance{}),
		Columns:     []schema.Column{},
	}
}
