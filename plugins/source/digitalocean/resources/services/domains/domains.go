package domains

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_domains",
		Resolver:  fetchDomainsDomains,
		Transform: transformers.TransformWithStruct(&godo.Domain{}),
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
			records(),
		},
	}
}
