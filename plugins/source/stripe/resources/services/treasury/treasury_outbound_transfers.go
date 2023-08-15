package treasury

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TreasuryOutboundTransfers() *schema.Table {
	return &schema.Table{
		Name:        "stripe_treasury_outbound_transfers",
		Description: `https://stripe.com/docs/api/treasury/outbound_transfers`,
		Transform:   client.TransformWithStruct(&stripe.TreasuryOutboundTransfer{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTreasuryOutboundTransfers,

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

func fetchTreasuryOutboundTransfers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	p := parent.Item.(*stripe.TreasuryFinancialAccount)

	lp := &stripe.TreasuryOutboundTransferListParams{
		FinancialAccount: stripe.String(p.ID),
	}

	it := cl.Services.TreasuryOutboundTransfers.List(lp)
	for it.Next() {
		res <- it.TreasuryOutboundTransfer()
	}

	return it.Err()
}
