package vendors

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchVendors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListVendorsWithContext(ctx, pagerduty.ListVendorOptions{
			Limit:  client.MaxPaginationLimit,
			Offset: offset,
		})
		if err != nil {
			return err
		}

		if len(response.Vendors) == 0 {
			return nil
		}

		res <- response.Vendors

		offset += uint(len(response.Vendors))
		more = response.More
	}

	return nil
}
