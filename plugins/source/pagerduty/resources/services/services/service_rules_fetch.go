package services

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchServiceRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)
	concreteParent := parent.Item.(pagerduty.Service)

	response, err := cqClient.PagerdutyClient.ListServiceRulesPaginated(ctx, concreteParent.ID)
	if err != nil {
		return err
	}

	res <- response

	return nil
}
