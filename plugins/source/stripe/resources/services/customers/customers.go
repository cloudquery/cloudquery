package customers

import (
	"context"

	"fmt"
	"strconv"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Customers() *schema.Table {
	return &schema.Table{
		Name:        "stripe_customers",
		Description: `https://stripe.com/docs/api/customers`,
		Transform:   transformers.TransformWithStruct(&stripe.Customer{}, transformers.WithSkipFields("APIResource", "ID"), transformers.WithIgnoreInTestsTransformer(client.CreateIgnoreInTestsTransformer("DefaultSource"))),
		Resolver:    fetchCustomers("customers"),

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
		IsIncremental: true,
	}
}

func fetchCustomers(tableName string) schema.TableResolver {
	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		cl := meta.(*client.Client)

		lp := &stripe.CustomerListParams{}

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

		it := cl.Services.Customers.List(lp)
		for it.Next() {
			data := it.Customer()
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
