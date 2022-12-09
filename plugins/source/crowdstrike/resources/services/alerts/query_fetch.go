package alerts

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/crowdstrike/client"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/crowdstrike/gofalcon/falcon/client/alerts"
)

func fetchQuery(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	desc := "timestamp.desc"
	retval, err := c.Services.Alerts.GetQueriesAlertsV1(&alerts.GetQueriesAlertsV1Params{
		Context: ctx,
		Sort:    &desc,
	})
	if err != nil {
		return err
	}
	payload := retval.GetPayload()
	res <- payload
	return nil
}
