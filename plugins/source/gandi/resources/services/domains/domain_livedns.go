package domains

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/go-gandi/go-gandi/domain"
)

func DomainLiveDNS() *schema.Table {
	return &schema.Table{
		Name:      "gandi_domain_livedns",
		Resolver:  fetchDomainLiveDNS,
		Transform: transformers.TransformWithStruct(&domain.LiveDNS{}),
		Columns: []schema.Column{
			{
				Name:       "fqdn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("fqdn"),
				PrimaryKey: true,
			},
		},
	}
}
