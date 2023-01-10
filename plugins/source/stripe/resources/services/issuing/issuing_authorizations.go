package issuing

import (
	"context"

	"fmt"
	"strconv"

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
		Resolver:    fetchIssuingAuthorizations("issuing_authorizations"),

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

func fetchIssuingAuthorizations(tableName string) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		cl := meta.(*client.Client)

		lp := &stripe.IssuingAuthorizationListParams{}

		if cl.Backend != nil {
			value, err := cl.Backend.Get(ctx, tableName, cl.ID())
			if err != nil {
				return fmt.Errorf("failed to retrieve state from backend: %w", err)
			}
			if value != "" {
				vi, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return fmt.Errorf("retrieved invalid state backend: %q %w", value, err)
				}
				lp.Created = &vi
			}
		}

		it := cl.Services.IssuingAuthorizations.List(lp)
		for it.Next() {
			res <- it.IssuingAuthorization()
		}
		return it.Err()
	}
}
