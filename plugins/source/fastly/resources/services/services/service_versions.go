package services

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/fastly/go-fastly/v7/fastly"
)

func ServiceVersions() *schema.Table {
	return &schema.Table{
		Name:        "fastly_service_versions",
		Description: `https://developer.fastly.com/reference/api/services/version/`,
		Resolver:    fetchServiceVersions,
		Transform:   transformers.TransformWithStruct(&fastly.Version{}),
		Columns: []schema.Column{
			{
				Name:       "number",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("Number"),
				PrimaryKey: true,
			},
			{
				Name:       "service_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ServiceID"),
				PrimaryKey: true,
			},
		},

		Relations: []*schema.Table{
			ServiceHealthChecks(),
			ServiceDomains(),
			ServiceBackends(),
		},
	}
}
