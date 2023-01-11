package terminal

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TerminalLocations() *schema.Table {
	return &schema.Table{
		Name:        "stripe_terminal_locations",
		Description: `https://stripe.com/docs/api/terminal_locations`,
		Transform:   transformers.TransformWithStruct(&stripe.TerminalLocation{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchTerminalLocations,

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

func fetchTerminalLocations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.TerminalLocationListParams{}

	it := cl.Services.TerminalLocations.List(lp)
	for it.Next() {
		res <- it.TerminalLocation()
	}

	return it.Err()
}
