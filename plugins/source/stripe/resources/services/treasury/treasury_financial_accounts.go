package treasury

import (
	"context"

	"fmt"
	"strconv"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryFinancialAccounts() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_financial_accounts",
		Description: `https://stripe.com/docs/api/treasury_financial_accounts`,
		Transform:   transformers.TransformWithStruct(&stripe.TreasuryFinancialAccount{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchTreasuryFinancialAccounts("treasury_financial_accounts"),

		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
				CreationOptions: schema.ColumnCreationOptions{
					IncrementalKey: true,
				},
			},
		},
		IsIncremental: true,

		Relations: []*schema.Table{
			TreasuryCreditReversals(),
			TreasuryDebitReversals(),
			TreasuryInboundTransfers(),
			TreasuryOutboundPayments(),
			TreasuryOutboundTransfers(),
			TreasuryReceivedCredits(),
			TreasuryReceivedDebits(),
			TreasuryTransactionEntries(),
			TreasuryTransactions(),
		},
	}
}

func fetchTreasuryFinancialAccounts(tableName string) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		cl := meta.(*client.Client)

		lp := &stripe.TreasuryFinancialAccountListParams{}

		if cl.Backend != nil {
			value, err := cl.Backend.Get(ctx, tableName, cl.ID())
			if err != nil {
				return fmt.Errorf("failed to retrieve state from backend: %w", err)
			}
			if value != "" {
				vi, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return fmt.Errorf("retrieved invalid state value: %q %w", value, err)
				}
				lp.Created = &vi
			}
		}

		it := cl.Services.TreasuryFinancialAccounts.List(lp)
		for it.Next() {
			data := it.TreasuryFinancialAccount()
			lp.Created = client.MaxInt64(lp.Created, &data.Created)
			res <- data
		}

		err := it.Err()
		if cl.Backend != nil && err == nil && lp.Created != nil {
			return cl.Backend.Set(ctx, tableName, cl.ID(), strconv.FormatInt(*lp.Created, 10))
		}
		return err
	}
}
