package teams

import (
	"context"
	"strconv"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/go-github/v49/github"
)

func members() *schema.Table {
	return &schema.Table{
		Name:      "github_team_members",
		Resolver:  fetchMembers,
		Transform: client.TransformWithStruct(&github.User{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			client.OrgColumn,
			teamIDColumn,
			{
				Name:     "membership",
				Type:     types.ExtensionTypes.JSON,
				Resolver: resolveMembership,
			},
		},
	}
}

func fetchMembers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
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

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
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
