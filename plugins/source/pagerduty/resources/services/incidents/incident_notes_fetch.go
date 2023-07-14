package incidents

import (
	"context"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchIncidentNotes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cqClient := meta.(*client.Client)
	concreteParent := parent.Item.(pagerduty.Incident)

	response, err := cqClient.PagerdutyClient.ListIncidentNotesWithContext(ctx, concreteParent.ID)
	if err != nil {
		return err
	}

	res <- response

	return nil
}
