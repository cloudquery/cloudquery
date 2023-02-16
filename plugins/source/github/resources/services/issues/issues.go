package issues

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/google/go-github/v48/github"
)

func Issues() *schema.Table {
	return &schema.Table{
		Name:      "github_issues",
		Resolver:  fetchIssues,
		Multiplex: client.OrgMultiplex,
		Transform: transformers.TransformWithStruct(&github.Issue{},
			append(client.SharedTransformers(), transformers.WithPrimaryKeys("ID"))...),
		Columns: []schema.Column{client.OrgColumn},
	}
}

func fetchIssues(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.IssueListOptions{
		Filter:      "all",
		State:       "all",
		ListOptions: github.ListOptions{PerPage: 100},
	}
	for {
		issues, resp, err := c.Github.Issues.ListByOrg(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		res <- issues
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
