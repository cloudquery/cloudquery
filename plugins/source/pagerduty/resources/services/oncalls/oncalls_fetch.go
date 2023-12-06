package oncalls

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/v4/schema"

	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
)

func fetchOncalls(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListOnCallsWithContext(ctx, pagerduty.ListOnCallOptions{
			Limit:    client.MaxPaginationLimit,
			Offset:   offset,
			TimeZone: "UTC",
		})
		if err != nil {
			return err
		}

		if len(response.OnCalls) == 0 {
			return nil
		}

		res <- response.OnCalls

		offset += uint(len(response.OnCalls))
		more = response.More
	}

	return nil
}
