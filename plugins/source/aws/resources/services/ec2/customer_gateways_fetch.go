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

func fetchEc2CustomerGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	response, err := svc.DescribeCustomerGateways(ctx, nil, func(options *ec2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- response.CustomerGateways
	return nil
}

func resolveCustomerGatewayArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	cg := resource.Item.(types.CustomerGateway)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "customer-gateway/" + aws.ToString(cg.CustomerGatewayId),
	}
	return resource.Set(c.Name, a.String())
}
