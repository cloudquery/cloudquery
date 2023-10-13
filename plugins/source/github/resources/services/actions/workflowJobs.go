package actions

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/google/go-github/v49/github"
)

func WorkflowJobs() *schema.Table {
	return &schema.Table{
		Name:      "github_workflow_jobs",
		Resolver:  fetchWorkflowJobs,
		Transform: client.TransformWithStruct(&github.WorkflowJob{}, transformers.WithPrimaryKeys("ID")),
		Columns: []schema.Column{
			client.OrgColumn,
			client.RepositoryIDColumn,
			{
				Name:        "run_id",
				Type:        arrow.PrimitiveTypes.Int64,
				Resolver:    schema.ParentColumnResolver("id"),
				Description: `Run ID`,
			},
		},
	}
}

func fetchWorkflowJobs(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := c.Repository
	run := resource.Item.(*github.WorkflowRun)

	listOpts := &github.ListOptions{PerPage: 100}
	listWorkflowJobOptions := &github.ListWorkflowJobsOptions{Filter: "all", ListOptions: *listOpts}
	for {
		workflowJobs, resp, err := c.Github.Actions.ListWorkflowJobs(ctx, c.Org, *repo.Name, *run.ID, listWorkflowJobOptions)
		if err != nil {
			return err
		}
		res <- workflowJobs.Jobs

		if resp.NextPage == 0 {
			break
		}
		listWorkflowJobOptions.Page = resp.NextPage
	}
	return nil
}
