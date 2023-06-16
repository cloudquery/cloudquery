package domains

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/go-gandi/go-gandi/domain"
)

func DomainWebRedirections() *schema.Table {
	return &schema.Table{
		Name:      "gandi_domain_web_redirections",
		Resolver:  fetchDomainWebRedirections,
		Transform: transformers.TransformWithStruct(&domain.WebRedirection{}),
		Columns: []schema.Column{
			{
				Name:       "fqdn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("fqdn"),
				PrimaryKey: true,
			},
			{
				Name:       "host",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Host"),
				PrimaryKey: true,
			},
			{
				Name:       "type",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Type"),
				PrimaryKey: true,
			},
		},
	}
}
