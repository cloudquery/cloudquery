package issuing

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func IssuingAuthorizations() *schema.Table {
	return &schema.Table{
		Name:        "stripe_issuing_authorizations",
		Description: `https://stripe.com/docs/api/issuing_authorizations`,
		Transform:   transformers.TransformWithStruct(&stripe.IssuingAuthorization{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchIssuingAuthorizations,

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

func fetchIssuingAuthorizations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.IssuingAuthorizations.List(&stripe.IssuingAuthorizationListParams{})
	for it.Next() {
		res <- it.IssuingAuthorization()
	}
	return it.Err()
}
