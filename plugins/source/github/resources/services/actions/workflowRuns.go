package actions

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func WorkflowRuns() *schema.Table {
	return &schema.Table{
		Name:      "github_workflow_runs",
		Resolver:  fetchWorkflowRuns,
		Multiplex: client.OrgRepositoryMultiplex,
		Transform: client.TransformWithStruct(&github.WorkflowRun{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			client.OrgColumn,
			client.RepositoryIDColumn,
		},
		Relations: []*schema.Table{
			WorkflowRunUsage(),
			WorkflowJobs(),
		},
	}
}

func fetchWorkflowRuns(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := c.Repository
	listOpts := &github.ListOptions{PerPage: 100}
	actionOpts := &github.ListWorkflowRunsOptions{ListOptions: *listOpts}
	for {
		workflowRuns, resp, err := c.Github.Actions.ListRepositoryWorkflowRuns(ctx, *repo.Owner.Login, *repo.Name, actionOpts)
		if err != nil {
			return err
		}
		res <- workflowRuns.WorkflowRuns

		if resp.NextPage == 0 {
			break
		}
		actionOpts.Page = resp.NextPage
	}
	return nil
}
