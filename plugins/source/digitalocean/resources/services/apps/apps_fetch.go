package apps

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/digitalocean/godo"
)

func fetchApps(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client)

	opt := godo.ListOptions{
		Page:    1,
		PerPage: client.MaxItemsPerPage,
	}
	return client.ThrottleWrapper(ctx, svc, func() error {
		for {
			// Fetch the current page and share the results.
			apps, resp, err := svc.Services.Apps.List(ctx, &opt)
			if err != nil {
				return err
			}
			res <- apps

			// Check if we're done.
			if resp.Links == nil || resp.Links.IsLastPage() {
				return nil
			}

			// Move onto the next page.
			page, err := resp.Links.CurrentPage()
			if err != nil {
				return fmt.Errorf("parsing response page number: %w", err)
			}
			opt.Page = page + 1
		}
	})
}
