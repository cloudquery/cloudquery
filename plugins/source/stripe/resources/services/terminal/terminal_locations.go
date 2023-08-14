package terminal

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TerminalLocations() *schema.Table {
	return &schema.Table{
		Name:        "stripe_terminal_locations",
		Description: `https://stripe.com/docs/api/terminal/locations`,
		Transform:   client.TransformWithStruct(&stripe.TerminalLocation{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTerminalLocations,

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

func fetchTerminalLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.TerminalLocationListParams{}

	it := cl.Services.TerminalLocations.List(lp)
	for it.Next() {
		res <- it.TerminalLocation()
	}

	return it.Err()
}
