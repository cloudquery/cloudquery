package services

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Services() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_services",
		Description: `https://developer.pagerduty.com/api-reference/e960cca205c0f-list-services`,
		Resolver:    fetchServices,
		Transform:   transformers.TransformWithStruct(&pagerduty.Service{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL")),
		Columns: []schema.Column{
			{
				Name:     "dependencies",
				Type:     schema.TypeJSON,
				Resolver: DependenciesResolver,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateAt"),
			},
			{
				Name:     "last_incident_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastIncidentTimestamp"),
			},
		},

		Relations: []*schema.Table{
			ServiceRules(),
		},
	}
}
