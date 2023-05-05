package cdns

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func Cdns() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_cdns",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#tag/CDN-Endpoints",
		Resolver:    fetchCdnsCdns,
		Transform:   transformers.TransformWithStruct(&godo.CDN{}),
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
