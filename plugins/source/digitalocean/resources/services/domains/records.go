package domains

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/digitalocean/godo"
)

func records() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_domain_records",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/domains_list_records",
		Resolver:    fetchDomainsRecords,
		Transform:   transformers.TransformWithStruct(&godo.DomainRecord{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
