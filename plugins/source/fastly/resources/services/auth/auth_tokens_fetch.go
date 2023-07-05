package auth

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func fetchAuthTokens(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		r, err := c.Fastly.ListTokens()
		if err != nil {
			return err
		}
		res <- r
		return nil
	}
	return c.RetryOnError(ctx, "fastly_auth_tokens", f)
}
