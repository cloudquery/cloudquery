package organizations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
)

func fetchMembers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.ListMembersOptions{ListOptions: github.ListOptions{PerPage: 100}}
	var orgMembers []*github.User
	for {
		members, resp, err := c.Github.Organizations.ListMembers(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- members
		orgMembers = append(orgMembers, members...)
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			return nil
		}
	}
}
