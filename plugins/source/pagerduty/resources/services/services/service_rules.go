package services

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ServiceRules() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_service_rules",
		Description: `https://developer.pagerduty.com/api-reference/d69ad7f1ec900-list-service-s-event-rules`,
		Resolver:    fetchServiceRules,
		Transform:   transformers.TransformWithStruct(&pagerduty.ServiceRule{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
