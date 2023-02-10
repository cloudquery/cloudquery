package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudformationStackSets(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	var config cloudformation.ListStackSetsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	paginator := cloudformation.NewListStackSetsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Summaries
	}
	return nil
}

func getStackSet(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	stack := resource.Item.(types.StackSetSummary)
	stackSet, err := meta.(*client.Client).Services().Cloudformation.DescribeStackSet(ctx, &cloudformation.DescribeStackSetInput{
		StackSetName: stack.StackSetName,
	})
	if err != nil {
		return err
	}
	resource.Item = stackSet.StackSet
	return nil
}

func fetchCloudformationStackSetOperations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(*types.StackSet)
	config := cloudformation.ListStackSetOperationsInput{
		StackSetName: stack.StackSetName,
	}
	svc := meta.(*client.Client).Services().Cloudformation
	paginator := cloudformation.NewListStackSetOperationsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Summaries
	}
	return nil
}

func getStackSetOperation(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	stack := resource.Parent.Item.(*types.StackSet)
	operation := resource.Item.(types.StackSetOperationSummary)
	stackSetOperation, err := meta.(*client.Client).Services().Cloudformation.DescribeStackSetOperation(ctx, &cloudformation.DescribeStackSetOperationInput{
		StackSetName: stack.StackSetName,
		OperationId:  operation.OperationId,
	})
	if err != nil {
		return err
	}
	resource.Item = stackSetOperation.StackSetOperation
	return nil
}

func fetchCloudformationStackSetOperationResults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stackSet := parent.Parent.Item.(*types.StackSet)
	operation := parent.Item.(*types.StackSetOperation)
	config := cloudformation.ListStackSetOperationResultsInput{
		OperationId:  operation.OperationId,
		StackSetName: stackSet.StackSetName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	paginator := cloudformation.NewListStackSetOperationResultsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Summaries
	}
	return nil
}
