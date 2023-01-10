package disputes

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Disputes() *schema.Table {
	return &schema.Table{
		Name:        "stripe_disputes",
		Description: `https://stripe.com/docs/api/disputes`,
		Transform:   transformers.TransformWithStruct(&stripe.Dispute{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchDisputes,

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
	}
}

func fetchDisputes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Disputes.List(&stripe.DisputeListParams{})
	for it.Next() {
		res <- it.Dispute()
	}
	return it.Err()
}
