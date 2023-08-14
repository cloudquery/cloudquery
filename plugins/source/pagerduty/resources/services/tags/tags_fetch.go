package tags

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchTags(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	response, err := cqClient.PagerdutyClient.ListTagsPaginated(ctx, pagerduty.ListTagOptions{
		Limit: client.MaxPaginationLimit,
	})
	if err != nil {
		return err
	}

	res <- response

	return nil
}
