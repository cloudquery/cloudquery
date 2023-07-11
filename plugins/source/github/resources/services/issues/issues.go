package issues

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func Issues() *schema.Table {
	return &schema.Table{
		Name:      "github_issues",
		Resolver:  fetchIssues,
		Multiplex: client.OrgRepositoryMultiplex,
		Transform: client.TransformWithStruct(&github.Issue{}, transformers.WithPrimaryKeys("ID")),
		Columns:   []schema.Column{client.OrgColumn, client.RepositoryIDColumn},
	}
}

func fetchIssues(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	opts := &github.IssueListByRepoOptions{
		State:       "all",
		ListOptions: github.ListOptions{PerPage: 100},
	}
	for {
		issues, resp, err := c.Github.Issues.ListByRepo(ctx, c.Org, *c.Repository.Name, opts)
		if err != nil {
			return err
		}
		res <- issues

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}
	return nil
}
