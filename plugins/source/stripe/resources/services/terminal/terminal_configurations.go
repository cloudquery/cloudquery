package terminal

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TerminalConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "stripe_terminal_configurations",
		Description: `https://stripe.com/docs/api/terminal_configurations`,
		Transform:   transformers.TransformWithStruct(&stripe.TerminalConfiguration{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchTerminalConfigurations,

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

func fetchTerminalConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.TerminalConfigurationListParams{}

	it := cl.Services.TerminalConfigurations.List(lp)
	for it.Next() {
		res <- it.TerminalConfiguration()
	}

	return it.Err()
}
