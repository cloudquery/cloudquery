package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/fastly/go-fastly/v7/fastly"
)

func fetchServiceBackends(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		v := parent.Item.(*fastly.Version)

		backend, err := c.Fastly.ListBackends(&fastly.ListBackendsInput{
			ServiceID:      v.ServiceID,
			ServiceVersion: v.Number,
		})
		if err != nil {
			return err
		}
		res <- backend
		return nil
	}
	return c.RetryOnError(ctx, "fastly_service_backends", f)
}
