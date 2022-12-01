package teams

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v48/github"
)

func fetchMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	t := parent.Item.(*github.Team)
	c := meta.(*client.Client)
	opts := &github.TeamListTeamMembersOptions{ListOptions: github.ListOptions{PerPage: 100}}
	orgId, err := strconv.Atoi(strings.Split(*t.MembersURL, "/")[4])
	if err != nil {
		return err
	}
	for {
		members, resp, err := c.Github.Teams.ListTeamMembersByID(ctx, int64(orgId), *t.ID, opts)
		if err != nil {
			return err
		}
		res <- members
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}

func resolveMembership(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, col schema.Column) error {
	c := meta.(*client.Client)

	m := resource.Item.(*github.User)
	t := resource.Parent.Item.(*github.Team)
	membership, _, err := c.Github.Teams.GetTeamMembershipBySlug(ctx, c.Org, *t.Slug, *m.Login)
	if err != nil {
		return err
	}
	return resource.Set(col.Name, membership)
}
