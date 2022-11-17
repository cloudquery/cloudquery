package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSsmInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm

	var input ssm.DescribeInstanceInformationInput
	for {
		output, err := svc.DescribeInstanceInformation(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.InstanceInformationList
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}

func fetchSsmInstanceComplianceItems(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(types.InstanceInformation)
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm

	input := ssm.ListComplianceItemsInput{
		ResourceIds: []string{*instance.InstanceId},
	}
	for {
		output, err := svc.ListComplianceItems(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.ComplianceItems
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}

func resolveInstanceARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Item.(types.InstanceInformation)
	cl := meta.(*client.Client)
	return resource.Set(c.Name, arn.ARN{
		Partition: cl.Partition,
		Service:   "ssm",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  fmt.Sprintf("managed-instance/%s", aws.ToString(instance.InstanceId)),
	}.String())
}
