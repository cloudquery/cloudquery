package balance

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Balance() *schema.Table {
	return &schema.Table{
		Name:        "stripe_balance",
		Description: `https://stripe.com/docs/api/balance`,
		Transform:   client.TransformWithStruct(&stripe.Balance{}, transformers.WithSkipFields("APIResource")),
		Resolver:    fetchBalance,
	}
}

func fetchBalance(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	val, err := cl.Services.Balance.Get(&stripe.BalanceParams{})
	if err != nil {
		return err
	}
	res <- val
	return nil
}
