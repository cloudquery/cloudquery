package credit_notes

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func CreditNotes() *schema.Table {
	return &schema.Table{
		Name:        "stripe_credit_notes",
		Description: `https://stripe.com/docs/api/credit_notes`,
		Transform:   transformers.TransformWithStruct(&stripe.CreditNote{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchCreditNotes,

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

func fetchCreditNotes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.CreditNoteListParams{}

	it := cl.Services.CreditNotes.List(lp)
	for it.Next() {
		res <- it.CreditNote()
	}

	return it.Err()
}
