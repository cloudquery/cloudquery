package images

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_images",
		Resolver:  fetchImagesImages,
		Transform: transformers.TransformWithStruct(&godo.Image{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
