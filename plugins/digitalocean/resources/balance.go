package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-digitalocean/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Balance() *schema.Table {
	return &schema.Table{
		Name:         "digitalocean_balance",
		Description:  "Balance represents a DigitalOcean Balance",
		Resolver:     fetchBalances,
		DeleteFilter: client.DeleteFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"generated_at"}},
		Columns: []schema.Column{
			{
				Name:        "month_to_date_balance",
				Description: "Balance as of the `generated_at` time.  This value includes the `account_balance` and `month_to_date_usage`.",
				Type:        schema.TypeString,
			},
			{
				Name:        "account_balance",
				Description: "Current balance of the customer's most recent billing activity.  Does not reflect `month_to_date_usage`.",
				Type:        schema.TypeString,
			},
			{
				Name:        "month_to_date_usage",
				Description: "Amount used in the current billing period as of the `generated_at` time.",
				Type:        schema.TypeString,
			},
			{
				Name:        "generated_at",
				Description: "The time at which balances were most recently generated.",
				Type:        schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchBalances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client)
	balance, _, err := svc.DoClient.Balance.Get(ctx)
	if err != nil {
		return err
	}
	if balance == nil {
		return nil
	}
	res <- *balance
	return nil
}
