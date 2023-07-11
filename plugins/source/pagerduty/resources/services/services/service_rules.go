package services

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ServiceRules() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_service_rules",
		Description: `https://developer.pagerduty.com/api-reference/d69ad7f1ec900-list-service-s-event-rules`,
		Resolver:    fetchServiceRules,
		Transform:   transformers.TransformWithStruct(&pagerduty.ServiceRule{}),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
		},
	}
}
