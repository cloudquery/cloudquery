package services

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "service_version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ServiceVersion"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
