// Code generated by codegen; DO NOT EDIT.

package files

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/stripe/stripe-go/v74"
)

func Files() *schema.Table {
	return &schema.Table{
		Name:        "stripe_files",
		Description: `https://stripe.com/docs/api/files`,
		Transform:   transformers.TransformWithStruct(&stripe.File{}, transformers.WithSkipFields("APIResource", "ID")),
		Resolver:    fetchFiles,

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

func fetchFiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	it := cl.Services.Files.List(&stripe.FileListParams{})
	for it.Next() {
		res <- it.File()
	}
	return it.Err()
}
