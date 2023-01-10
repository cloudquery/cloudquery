package file_links

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func FileLinks() *schema.Table {
	return &schema.Table{
		Name:        "stripe_file_links",
		Description: `https://stripe.com/docs/api/file_links`,
		Transform:   transformers.TransformWithStruct(&stripe.FileLink{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchFileLinks,

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

func fetchFileLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.FileLinks.List(&stripe.FileLinkListParams{})
	for it.Next() {
		res <- it.FileLink()
	}
	return it.Err()
}
