package monitoring

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Name:       "uuid",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("UUID"),
				PrimaryKey: true,
			},
		},
	}
}
