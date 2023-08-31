package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/fastly/go-fastly/v7/fastly"
)

func fetchServiceHealthChecks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		v := parent.Item.(*fastly.Version)
		healthChecks, err := c.Fastly.ListHealthChecks(&fastly.ListHealthChecksInput{
			ServiceID:      v.ServiceID,
			ServiceVersion: v.Number,
		})
		if err != nil {
			return err
		}
		res <- healthChecks

		return nil
	}
	return c.RetryOnError(ctx, "fastly_service_health_checks", f)
}
