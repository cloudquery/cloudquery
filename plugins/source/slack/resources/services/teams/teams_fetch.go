package teams

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchTeams(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	f := func() error {
		teams, err := c.Slack.GetTeamInfoContext(ctx)
		if err != nil {
			return err
		}
		res <- teams
		return nil
	}
	return c.RetryOnError(ctx, "slack_teams", f)
}
