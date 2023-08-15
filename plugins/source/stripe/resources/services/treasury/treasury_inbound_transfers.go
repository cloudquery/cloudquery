package treasury

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryInboundTransfers() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_inbound_transfers",
		Description: `https://stripe.com/docs/api/treasury/inbound_transfers`,
		Transform:   client.TransformWithStruct(&stripe.TreasuryInboundTransfer{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTreasuryInboundTransfers,

		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchTreasuryInboundTransfers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.TreasuryFinancialAccount)

	lp := &stripe.TreasuryInboundTransferListParams{
		FinancialAccount: stripe.String(p.ID),
	}

	it := cl.Services.TreasuryInboundTransfers.List(lp)
	for it.Next() {
		res <- it.TreasuryInboundTransfer()
	}

	return it.Err()
}
