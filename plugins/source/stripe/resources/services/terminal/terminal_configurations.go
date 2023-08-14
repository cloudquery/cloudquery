package terminal

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TerminalConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "stripe_terminal_configurations",
		Description: `https://stripe.com/docs/api/terminal/configuration`,
		Transform:   client.TransformWithStruct(&stripe.TerminalConfiguration{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTerminalConfigurations,

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

func fetchTerminalConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.TerminalConfigurationListParams{}

	it := cl.Services.TerminalConfigurations.List(lp)
	for it.Next() {
		res <- it.TerminalConfiguration()
	}

	return it.Err()
}
