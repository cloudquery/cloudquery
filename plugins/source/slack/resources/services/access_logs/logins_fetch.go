package access_logs

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/slack-go/slack"
)

func fetchLogins(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	params := slack.AccessLogParameters{
		Count:  1000,
		TeamID: meta.(*client.Client).TeamID,
		Page:   0,
	}
	slackClient := meta.(*client.Client).Slack
	for {
		logins, paging, err := slackClient.GetAccessLogsContext(ctx, params)
		if err != nil {
			return err
		}
		res <- logins
		if paging.Page >= paging.Pages {
			break
		}
		params.Page = paging.Page + 1
	}
	return nil
}
