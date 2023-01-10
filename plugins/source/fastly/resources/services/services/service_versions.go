package services

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
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
				Name:     "number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Number"),
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
		},

		Relations: []*schema.Table{
			ServiceHealthChecks(),
			ServiceDomains(),
			ServiceBackends(),
		},
	}
}
