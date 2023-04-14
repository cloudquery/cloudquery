package sizes

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func Sizes() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_sizes",
		Resolver:  fetchSizesSizes,
		Transform: transformers.TransformWithStruct(&godo.Size{}),
		Columns: []schema.Column{
			{
				Name:     "slug",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Slug"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
