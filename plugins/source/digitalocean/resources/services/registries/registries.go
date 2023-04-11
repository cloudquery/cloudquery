package registries

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func Registries() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_registries",
		Resolver:  fetchRegistriesRegistries,
		Transform: transformers.TransformWithStruct(&godo.Registry{}),
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			repositories(),
		},
	}
}
