package sizes

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func Sizes() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_sizes",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Sizes",
		Resolver:    fetchSizesSizes,
		Transform:   transformers.TransformWithStruct(&godo.Size{}),
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
