package domains

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/go-gandi/go-gandi/domain"
)

func DomainWebRedirections() *schema.Table {
	return &schema.Table{
		Name:      "gandi_domain_web_redirections",
		Resolver:  fetchDomainWebRedirections,
		Transform: transformers.TransformWithStruct(&domain.WebRedirection{}),
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
				Name:     "host",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Host"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
