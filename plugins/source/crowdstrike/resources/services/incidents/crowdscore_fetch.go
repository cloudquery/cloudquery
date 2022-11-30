package incidents

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/crowdstrike/client"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/crowdstrike/gofalcon/falcon/client/incidents"
)

func fetchCrowdscore(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	desc := "timestamp.desc"
	retval, err := c.Services.Incidents.CrowdScore(&incidents.CrowdScoreParams{
		Context: ctx,
		Sort:    &desc,
	})
	if err != nil {
		return err
	}
	payload := retval.GetPayload()
	resources := payload.Resources
	res <- resources
	return nil
}
