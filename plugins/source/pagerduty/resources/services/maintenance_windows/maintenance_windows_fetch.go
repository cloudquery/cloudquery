package maintenance_windows

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchMaintenanceWindows(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListMaintenanceWindowsWithContext(ctx, pagerduty.ListMaintenanceWindowsOptions{
			Limit:   client.MaxPaginationLimit,
			Offset:  offset,
			TeamIDs: cqClient.Spec.TeamIds,
		})
		if err != nil {
			return err
		}

		if len(response.MaintenanceWindows) == 0 {
			return nil
		}

		res <- response.MaintenanceWindows

		offset += uint(len(response.MaintenanceWindows))
		more = response.More
	}

	return nil
}
