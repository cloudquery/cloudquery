package images

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func Images() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_images",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/Images",
		Resolver:    fetchImagesImages,
		Transform:   transformers.TransformWithStruct(&godo.Image{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
