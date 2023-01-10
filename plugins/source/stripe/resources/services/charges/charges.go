package charges

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Charges() *schema.Table {
	return &schema.Table{
		Name:        "stripe_charges",
		Description: `https://stripe.com/docs/api/charges`,
		Transform:   transformers.TransformWithStruct(&stripe.Charge{}, transformers.WithSkipFields("APIResource", "ID"), transformers.WithIgnoreInTestsTransformer(client.CreateIgnoreInTestsTransformer("Destination", "Dispute", "Level3", "Source"))),
		Resolver:    fetchCharges,

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

func fetchCharges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Charges.List(&stripe.ChargeListParams{})
	for it.Next() {
		res <- it.Charge()
	}
	return it.Err()
}
