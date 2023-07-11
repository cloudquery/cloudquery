package organizations

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/google/go-github/v49/github"
)

func members() *schema.Table {
	return &schema.Table{
		Name:      "github_organization_members",
		Resolver:  fetchMembers,
		Transform: client.TransformWithStruct(&github.User{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			client.OrgColumn,
			{
				Name:     "membership",
				Type:     types.ExtensionTypes.JSON,
				Resolver: resolveMembership,
			},
		},
	}
}

func fetchMembers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.ListMembersOptions{ListOptions: github.ListOptions{PerPage: 100}}
	for {
		members, resp, err := c.Github.Organizations.ListMembers(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- members

		if resp.NextPage == 0 {
			return nil
		}
		opts.Page = resp.NextPage
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
