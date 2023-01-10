package accounts

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Accounts() *schema.Table {
	return &schema.Table{
		Name:        "stripe_accounts",
		Description: `https://stripe.com/docs/api/accounts`,
		Transform:   transformers.TransformWithStruct(&stripe.Account{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchAccounts,

		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			Capabilities(),
		},
	}
}

func fetchAccounts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Accounts.List(&stripe.AccountListParams{})
	for it.Next() {
		res <- it.Account()
	}
	return it.Err()
}
