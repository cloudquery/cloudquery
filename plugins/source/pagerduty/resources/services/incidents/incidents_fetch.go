package incidents

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchIncidents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListIncidentsWithContext(ctx, pagerduty.ListIncidentsOptions{
			Limit:     client.MaxPaginationLimit,
			Offset:    offset,
			DateRange: "all",
			TeamIDs:   cqClient.Spec.TeamIds,
		})
		if err != nil {
			return err
		}

		if len(response.Incidents) == 0 {
			return nil
		}

		res <- response.Incidents

		offset += uint(len(response.Incidents))
		more = response.More
	}

	return nil
}
