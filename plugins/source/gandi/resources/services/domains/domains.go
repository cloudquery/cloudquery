package domains

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "sharing_id",
				Type:     schema.TypeString,
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
