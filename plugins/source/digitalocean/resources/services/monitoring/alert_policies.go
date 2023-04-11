package monitoring

import (
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/digitalocean/godo"
)

func AlertPolicies() *schema.Table {
	return &schema.Table{
		Name:      "digitalocean_monitoring_alert_policies",
		Resolver:  fetchMonitoringAlertPolicies,
		Transform: transformers.TransformWithStruct(&godo.AlertPolicy{}),
		Columns: []schema.Column{
			{
				Name:     "uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UUID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
