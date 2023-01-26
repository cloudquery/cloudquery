package rulesets

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchRulesets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	response, err := cqClient.PagerdutyClient.ListRulesetsPaginated(ctx)
	if err != nil {
		return err
	}

	res <- response

	return nil
}
