package priorities

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchPriorities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListPrioritiesWithContext(ctx, pagerduty.ListPrioritiesOptions{
			Limit:  client.MaxPaginationLimit,
			Offset: offset,
		})
		if err != nil {
			return err
		}

		if len(response.Priorities) == 0 {
			return nil
		}

		res <- response.Priorities

		offset += uint(len(response.Priorities))
		more = response.More
	}

	return nil
}
