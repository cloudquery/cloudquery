package services

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
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
				Type:     types.ExtensionTypes.JSON,
				Resolver: DependenciesResolver,
			},
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "html_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "created_at",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("CreateAt"),
			},
			{
				Name:     "last_incident_timestamp",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("LastIncidentTimestamp"),
			},
		},

		Relations: []*schema.Table{
			ServiceRules(),
		},
	}
}
