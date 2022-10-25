package actions

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/google/go-github/v45/github"
)

func fetchWorkflows(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	opts := &github.RepositoryListByOrgOptions{ListOptions: github.ListOptions{PerPage: 100}}
	for {
		repos, resp, err := c.Github.Repositories.ListByOrg(ctx, c.Org, opts)
		if err != nil {
			return err
		}
		for _, repo := range repos {
			actionOpts := &github.ListOptions{PerPage: 100}
			for {
				workflows, resp, err := c.Github.Actions.ListWorkflows(ctx, *repo.Owner.Login, *repo.Name, actionOpts)
				if err != nil {
					return err
				}
				res <- workflows.Workflows
				opts.Page = resp.NextPage
				if opts.Page == resp.LastPage {
					break
				}
			}
		}
		opts.Page = resp.NextPage
		if opts.Page == resp.LastPage {
			break
		}
	}
	return nil
}
