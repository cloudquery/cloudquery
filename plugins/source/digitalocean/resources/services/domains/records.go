package domains

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
