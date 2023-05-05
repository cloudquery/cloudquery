package droplets

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			neighbors(),
		},
	}
}
