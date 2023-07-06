package services

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/fastly/go-fastly/v7/fastly"
)

func ServiceHealthChecks() *schema.Table {
	return &schema.Table{
		Name:        "fastly_service_health_checks",
		Description: `https://developer.fastly.com/reference/api/services/healthcheck/`,
		Resolver:    fetchServiceHealthChecks,
		Transform:   transformers.TransformWithStruct(&fastly.HealthCheck{}),
		Columns: []schema.Column{
			{
				Name:       "name",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("Name"),
				PrimaryKey: true,
			},
			{
				Name:       "service_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ServiceID"),
				PrimaryKey: true,
			},
			{
				Name:       "service_version",
				Type:       arrow.PrimitiveTypes.Int64,
				Resolver:   schema.PathResolver("ServiceVersion"),
				PrimaryKey: true,
			},
		},
	}
}
