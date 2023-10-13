package actions

import (
	"context"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/google/go-github/v49/github"
)

func workflowRunUsage() *schema.Table {
	return &schema.Table{
		Name:      "github_workflow_run_usage",
		Resolver:  fetchWorkflowRunUsage,
		Transform: client.TransformWithStruct(&github.WorkflowRunUsage{}),
		Columns: []schema.Column{
			client.OrgColumn,
			client.RepositoryIDColumn,
			{
				Name:        "run_id",
				Type:        arrow.PrimitiveTypes.Int64,
				Resolver:    schema.ParentColumnResolver("id"),
				Description: `Run ID`,
				PrimaryKey:  true,
			},
		},
	}
}

func fetchWorkflowRunUsage(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	repo := c.Repository
	run := resource.Item.(*github.WorkflowRun)

	workflowRunUsage, _, err := c.Github.Actions.GetWorkflowRunUsageByID(ctx, c.Org, *repo.Name, *run.ID)
	if err != nil {
		return err
	}

	res <- workflowRunUsage

	return nil
}
