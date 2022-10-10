package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlueWorkflows(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	input := glue.ListWorkflowsInput{MaxResults: aws.Int32(25)}
	for {
		result, err := svc.ListWorkflows(ctx, &input)
		if err != nil {
			return err
		}
		res <- result.Workflows

		if aws.ToString(result.NextToken) == "" {
			break
		}
		input.NextToken = result.NextToken
	}
	return nil
}

func getWorkflow(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	wf := resource.Item.(string)

	w, err := svc.GetWorkflow(ctx, &glue.GetWorkflowInput{Name: &wf})
	if err != nil {
		return err
	}

	resource.Item = w.Workflow
	return nil
}

func resolveGlueWorkflowArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	arn := aws.String(workflowARN(cl, aws.ToString(resource.Item.(*types.Workflow).Name)))
	return resource.Set(c.Name, arn)
}

func resolveGlueWorkflowTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetTags(ctx, &glue.GetTagsInput{
		ResourceArn: aws.String(workflowARN(cl, aws.ToString(resource.Item.(*types.Workflow).Name))),
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Tags)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func workflowARN(cl *client.Client, name string) string {
	return cl.ARN(client.GlueService, "workflow", name)
}
