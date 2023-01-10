package domains

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/go-gandi/go-gandi/domain"
)

func DomainDNSSecKeys() *schema.Table {
	return &schema.Table{
		Name:      "gandi_domain_dnssec_keys",
		Resolver:  fetchDomainDNSSecKeys,
		Transform: transformers.TransformWithStruct(&domain.DNSSECKey{}),
		Columns: []schema.Column{
			{
				Name:     "fqdn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("fqdn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
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
