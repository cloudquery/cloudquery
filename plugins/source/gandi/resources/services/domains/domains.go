package domains

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/go-gandi/go-gandi/domain"
)

func Domains() *schema.Table {
	return &schema.Table{
		Name:                "gandi_domains",
		Resolver:            fetchDomains,
		Transform:           transformers.TransformWithStruct(&domain.Details{}),
		PreResourceResolver: getDomain,
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "sharing_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("SharingID"),
			},
		},

		Relations: []*schema.Table{
			DomainLiveDNS(),
			DomainWebRedirections(),
			DomainGlueRecords(),
			DomainDNSSecKeys(),
		},
	}
}
