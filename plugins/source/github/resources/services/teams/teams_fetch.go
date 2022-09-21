package teams

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
)

func fetchTeams(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.ListOptions{PerPage: 100}
	for {
		repos, resp, err := c.Github.Teams.ListTeams(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- repos
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
