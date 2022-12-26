package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/fastly/go-fastly/v7/fastly"
)

func fetchServiceVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	s := parent.Item.(fastly.Service)
	versions, err := c.Fastly.ListVersions(&fastly.ListVersionsInput{
		ServiceID: s.ID,
	})
	if err != nil {
		return err
	}
	res <- versions
	return nil
}
