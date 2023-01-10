package droplets

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Droplets() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_droplets",
		Resolver:  fetchDropletsDroplets,
		Transform: transformers.TransformWithStruct(&godo.Droplet{}),
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
			Neighbors(),
		},
	}
}
