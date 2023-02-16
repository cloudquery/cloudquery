package external

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "github_external_groups",
		Resolver:  fetchGroups,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.ExternalGroup{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("GroupID"))...),
		Columns: []schema.Column{client.OrgColumn},
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
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
