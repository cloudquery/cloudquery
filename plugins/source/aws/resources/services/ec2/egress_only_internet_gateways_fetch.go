package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEc2EgressOnlyInternetGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	input := ec2.DescribeEgressOnlyInternetGatewaysInput{}
	for {
		output, err := svc.DescribeEgressOnlyInternetGateways(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.EgressOnlyInternetGateways
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}

func resolveEgressOnlyInternetGatewaysArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.EgressOnlyInternetGateway)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "egress-only-internet-gateway/" + aws.ToString(item.EgressOnlyInternetGatewayId),
	}
	return resource.Set(c.Name, a.String())
}
