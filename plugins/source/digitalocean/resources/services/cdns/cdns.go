package cdns

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Cdns() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_cdns",
		Resolver:  fetchCdnsCdns,
		Transform: transformers.TransformWithStruct(&godo.CDN{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
