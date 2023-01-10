package files

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchFiles(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		params := &slack.GetFilesParameters{
			Count:      1000,
			ShowHidden: true,
		}
		for {
			files, paging, err := c.Slack.GetFilesContext(ctx, *params)
			if err != nil {
				return err
			}
			res <- files
			if paging.Page >= paging.Pages || paging.Pages == 1 {
				break
			}
			params.Page = paging.Page + 1
		}
		return nil
	}
	return c.RetryOnError(ctx, "slack_files", f)
}
