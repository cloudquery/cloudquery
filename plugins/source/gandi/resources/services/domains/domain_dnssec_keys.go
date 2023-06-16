package domains

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/go-gandi/go-gandi/domain"
)

func DomainDNSSecKeys() *schema.Table {
	return &schema.Table{
		Name:      "gandi_domain_dnssec_keys",
		Resolver:  fetchDomainDNSSecKeys,
		Transform: transformers.TransformWithStruct(&domain.DNSSECKey{}),
		Columns: []schema.Column{
			{
				Name:       "fqdn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("fqdn"),
				PrimaryKey: true,
			},
			{
				Name:       "id",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
