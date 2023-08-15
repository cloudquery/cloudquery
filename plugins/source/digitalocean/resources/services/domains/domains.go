package domains

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_domains",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/domains_list",
		Resolver:    fetchDomainsDomains,
		Transform:   transformers.TransformWithStruct(&godo.Domain{}),
		Columns: []schema.Column{
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			records(),
		},
	}
}
