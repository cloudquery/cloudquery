package monitoring

import (
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
	"github.com/digitalocean/godo"
)

func AlertPolicies() *schema.Table {
	return &schema.Table{
		Name:        "digitalocean_monitoring_alert_policies",
		Description: "https://docs.digitalocean.com/reference/api/api-reference/#operation/monitoring_list_alertPolicy",
		Resolver:    fetchMonitoringAlertPolicies,
		Transform:   transformers.TransformWithStruct(&godo.AlertPolicy{}),
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
