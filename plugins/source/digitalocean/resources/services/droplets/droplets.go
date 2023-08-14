package droplets

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func Droplets() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_droplets",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/droplets_list",
		Resolver:    fetchDropletsDroplets,
		Transform:   transformers.TransformWithStruct(&godo.Droplet{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			neighbors(),
		},
	}
}
