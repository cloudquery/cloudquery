package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/fastly/go-fastly/v7/fastly"
)

func fetchServiceVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		s := parent.Item.(*fastly.Service)
		versions, err := c.Fastly.ListVersions(&fastly.ListVersionsInput{
			ServiceID: s.ID,
		})
		if err != nil {
			return err
		}
		res <- versions
		return nil
	}
	return c.RetryOnError(ctx, "fastly_service_versions", f)
}
