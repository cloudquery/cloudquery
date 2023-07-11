package extensions

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchExtensions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListExtensionsWithContext(ctx, pagerduty.ListExtensionOptions{
			Limit:  client.MaxPaginationLimit,
			Offset: offset,
		})
		if err != nil {
			return err
		}

		if len(response.Extensions) == 0 {
			return nil
		}

		res <- response.Extensions

		offset += uint(len(response.Extensions))
		more = response.More
	}

	return nil
}
