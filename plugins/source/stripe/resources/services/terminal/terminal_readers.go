package terminal

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TerminalReaders() *schema.Table {
	return &schema.Table{
		Name:        "stripe_terminal_readers",
		Description: `https://stripe.com/docs/api/terminal_readers`,
		Transform:   transformers.TransformWithStruct(&stripe.TerminalReader{}, client.SharedTransformers(transformers.WithSkipFields("APIResource", "ID"))...),
		Resolver:    fetchTerminalReaders,

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

func fetchTerminalReaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.TerminalReaderListParams{}

	it := cl.Services.TerminalReaders.List(lp)
	for it.Next() {
		res <- it.TerminalReader()
	}

	return it.Err()
}
