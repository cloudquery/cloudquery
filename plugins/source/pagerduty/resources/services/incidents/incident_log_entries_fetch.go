package incidents

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchIncidentLogEntries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)
	concreteParent := parent.Item.(pagerduty.Incident)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListIncidentLogEntriesWithContext(ctx, concreteParent.ID, pagerduty.ListIncidentLogEntriesOptions{
			Limit:  client.MaxPaginationLimit,
			Offset: offset,
		})
		if err != nil {
			return err
		}

		if len(response.LogEntries) == 0 {
			return nil
		}

		res <- response.LogEntries

		offset += uint(len(response.LogEntries))
		more = response.More
	}

	return nil
}
