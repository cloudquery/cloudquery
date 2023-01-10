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

func fetchEc2ReservedInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ec2.DescribeReservedInstancesInput
	c := meta.(*client.Client)
	svc := c.Services().Ec2

	// this API does not seem to support any form of pagination
	output, err := svc.DescribeReservedInstances(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.ReservedInstances
	return nil
}

func resolveReservedInstanceArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ReservedInstances)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "reserved-instance/" + aws.ToString(item.ReservedInstancesId),
	}
	return resource.Set(c.Name, a.String())
}
