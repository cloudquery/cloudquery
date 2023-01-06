package incidents

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func IncidentAlerts() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_incident_alerts",
		Description: `https://developer.pagerduty.com/api-reference/4bc42e7ac0c59-list-alerts-for-an-incident`,
		Resolver:    fetchIncidentAlerts,
		Transform:   transformers.TransformWithStruct(&pagerduty.IncidentAlert{}, transformers.WithSkipFields("HTMLURL")),
		Columns: []schema.Column{
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
		},
	}
}
