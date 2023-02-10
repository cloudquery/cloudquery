package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudformationStacks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	var config cloudformation.DescribeStacksInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	paginator := cloudformation.NewDescribeStacksPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.Stacks
	}
	return nil
}
func fetchCloudformationStackResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(types.Stack)
	config := cloudformation.ListStackResourcesInput{
		StackName: stack.StackName,
	}
	svc := meta.(*client.Client).Services().Cloudformation
	paginator := cloudformation.NewListStackResourcesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.StackResourceSummaries
	}
	return nil
}
