package issuing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func IssuingDisputes() *schema.Table {
	return &schema.Table{
		Name:        "stripe_issuing_disputes",
		Description: `https://stripe.com/docs/api/issuing_disputes`,
		Transform:   transformers.TransformWithStruct(&stripe.IssuingDispute{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchIssuingDisputes,

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

func fetchIssuingDisputes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.IssuingDisputes.List(&stripe.IssuingDisputeListParams{})
	for it.Next() {
		res <- it.IssuingDispute()
	}
	return it.Err()
}
