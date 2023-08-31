package external

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "github_external_groups",
		Resolver:  fetchGroups,
		Multiplex: client.OrgMultiplex,
		Transform: client.TransformWithStruct(&github.ExternalGroup{}, transformers.WithPrimaryKeys("GroupID")),
		Columns:   []schema.Column{client.OrgColumn},
	}
}

func fetchGroups(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.ListExternalGroupsOptions{ListOptions: github.ListOptions{PerPage: 100}}
	for {
		groups, resp, err := c.Github.Teams.ListExternalGroups(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- groups.Groups

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return nil
}
