package registries

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func repositories() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_registry_repositories",
		Resolver:  fetchRegistriesRepositories,
		Transform: transformers.TransformWithStruct(&godo.Repository{}),
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
	}
}
