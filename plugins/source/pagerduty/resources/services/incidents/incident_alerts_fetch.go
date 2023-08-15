package incidents

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchIncidentAlerts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)
	concreteParent := parent.Item.(pagerduty.Incident)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListIncidentAlertsWithContext(ctx, concreteParent.ID, pagerduty.ListIncidentAlertsOptions{
			Limit:  client.MaxPaginationLimit,
			Offset: offset,
		})
		if err != nil {
			return err
		}

		if len(response.Alerts) == 0 {
			return nil
		}

		res <- response.Alerts

		offset += uint(len(response.Alerts))
		more = response.More
	}

	return nil
}
