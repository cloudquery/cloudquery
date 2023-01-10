package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudformationStacks(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	var config cloudformation.DescribeStacksInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	for {
		output, err := svc.DescribeStacks(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.Stacks
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func fetchCloudformationStackResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(types.Stack)
	config := cloudformation.ListStackResourcesInput{
		StackName: stack.StackName,
	}
	c := meta.(*client.Client)
	svc := c.Services().Cloudformation
	for {
		output, err := svc.ListStackResources(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.StackResourceSummaries
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
