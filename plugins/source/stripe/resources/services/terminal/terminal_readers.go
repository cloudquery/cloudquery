package terminal

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/stripe/stripe-go/v74"
)

func TerminalReaders() *schema.Table {
	return &schema.Table{
		Name:        "stripe_terminal_readers",
		Description: `https://stripe.com/docs/api/terminal/readers`,
		Transform:   client.TransformWithStruct(&stripe.TerminalReader{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchTerminalReaders,

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

func fetchTerminalReaders(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	lp := &stripe.TerminalReaderListParams{}

	it := cl.Services.TerminalReaders.List(lp)
	for it.Next() {
		res <- it.TerminalReader()
	}

	return it.Err()
}
