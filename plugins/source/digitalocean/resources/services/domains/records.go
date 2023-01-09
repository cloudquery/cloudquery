package domains

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/digitalocean/godo"
)

func Records() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_domain_records",
		Resolver:  fetchDomainsRecords,
		Transform: transformers.TransformWithStruct(&godo.DomainRecord{}),
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
