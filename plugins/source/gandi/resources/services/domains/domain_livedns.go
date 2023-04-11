package domains

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/go-gandi/go-gandi/domain"
)

func DomainLiveDNS() *schema.Table {
	return &schema.Table{
		Name:      "gandi_domain_livedns",
		Resolver:  fetchDomainLiveDNS,
		Transform: transformers.TransformWithStruct(&domain.LiveDNS{}),
		Columns: []schema.Column{
			{
				Name:     "fqdn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("fqdn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
