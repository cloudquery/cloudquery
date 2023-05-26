package domains

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/go-gandi/go-gandi/domain"
)

func DomainGlueRecords() *schema.Table {
	return &schema.Table{
		Name:      "gandi_domain_glue_records",
		Resolver:  fetchDomainGlueRecords,
		Transform: transformers.TransformWithStruct(&domain.GlueRecord{}),
		Columns: []schema.Column{
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
			{
				Name:       "fqdn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("FQDN"),
				PrimaryKey: true,
			},
		},
	}
}
