package escalation_policies

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchEscalationPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)

	more := true
	var offset uint
	for more {
		response, err := cqClient.PagerdutyClient.ListEscalationPoliciesWithContext(ctx, pagerduty.ListEscalationPoliciesOptions{
			Limit:   client.MaxPaginationLimit,
			Offset:  offset,
			TeamIDs: cqClient.Spec.TeamIds,
		})
		if err != nil {
			return err
		}

		if len(response.EscalationPolicies) == 0 {
			return nil
		}

		res <- response.EscalationPolicies

		offset += uint(len(response.EscalationPolicies))
		more = response.More
	}

	return nil
}
