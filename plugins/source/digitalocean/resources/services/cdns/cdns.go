package cdns

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
