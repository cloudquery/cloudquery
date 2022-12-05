package organizations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func fetchMembers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.ListMembersOptions{ListOptions: github.ListOptions{PerPage: 100}}
	for {
		members, resp, err := c.Github.Organizations.ListMembers(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- members
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			return nil
		}
	}
}

func resolveMembership(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	c := meta.(*client.Client)

	m := resource.Item.(*github.User)
	membership, _, err := c.Github.Organizations.GetOrgMembership(ctx, *m.Login, c.Org)
	if err != nil {
		return err
	}
	return resource.Set(col.Name, membership)
}
