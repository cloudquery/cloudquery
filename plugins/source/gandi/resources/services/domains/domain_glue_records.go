package domains

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/go-gandi/go-gandi/domain"
)

func DomainGlueRecords() *schema.Table {
	return &schema.Table{
		Name:      "gandi_domain_glue_records",
		Resolver:  fetchDomainGlueRecords,
		Transform: transformers.TransformWithStruct(&domain.GlueRecord{}),
		Columns: []schema.Column{
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "fqdn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FQDN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
