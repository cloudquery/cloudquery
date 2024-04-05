package actions

import (
	"context"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v59/github"
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
			workflowRunUsage(),
			workflowJobs(),
		},
	}
}

func fetchWorkflowRuns(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := c.Repository
	listOpts := &github.ListOptions{PerPage: 100}
	actionOpts := &github.ListWorkflowRunsOptions{ListOptions: *listOpts}

	if c.Spec.TableOptions.WorkflowRuns.ParsedTimeSince != "" {
		actionOpts.Created = ">=" + c.Spec.TableOptions.WorkflowRuns.ParsedTimeSince
	}
	var lastCreatedAt string
	for {
		workflowRuns, resp, err := c.Github.Actions.ListRepositoryWorkflowRuns(ctx, *repo.Owner.Login, *repo.Name, actionOpts)
		if err != nil {
			return err
		}
		// When setting created_at, the API will return a maximum of 1000 records, so if we reached the limit we need to get the other records
		// Workflows runs are sorted by created_at in descending order, so we need to scope down the query
		if len(workflowRuns.WorkflowRuns) == 0 && actionOpts.Page > 0 && actionOpts.Created != "" {
			actionOpts.Page = 0
			actionOpts.Created = c.Spec.TableOptions.WorkflowRuns.ParsedTimeSince + ".." + lastCreatedAt
			workflowRuns, resp, err = c.Github.Actions.ListRepositoryWorkflowRuns(ctx, *repo.Owner.Login, *repo.Name, actionOpts)
			if err != nil {
				return err
			}
		}
		if len(workflowRuns.WorkflowRuns) > 0 {
			lastCreatedAt = workflowRuns.WorkflowRuns[len(workflowRuns.WorkflowRuns)-1].GetCreatedAt().Format(time.RFC3339)
		}
		res <- workflowRuns.WorkflowRuns

		if resp.NextPage == 0 {
			break
		}
		actionOpts.Page = resp.NextPage
	}
	return nil
}
