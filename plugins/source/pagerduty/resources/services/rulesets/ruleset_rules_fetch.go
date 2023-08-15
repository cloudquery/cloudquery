package rulesets

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchRulesetRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)
	concreteParent := parent.Item.(*pagerduty.Ruleset)

	response, err := cqClient.PagerdutyClient.ListRulesetRulesPaginated(ctx, concreteParent.ID)
	if err != nil {
		return err
	}

	res <- response

	return nil
}
