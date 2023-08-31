package account

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/fastly/go-fastly/v7/fastly"
)

func fetchAccountEvents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		input := &fastly.GetAPIEventsFilterInput{
			MaxResults: 100,
			PageNumber: 1,
		}
		for {
			r, err := c.Fastly.GetAPIEvents(input)
			if err != nil {
				return err
			}
			res <- r.Events
			if r.Links.Next == "" {
				break
			}
			input.PageNumber++
		}
		return nil
	}
	return c.RetryOnError(ctx, "fastly_account_events", f)
}
