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

func fetchEc2VpnGateways(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ec2.DescribeVpnGatewaysInput
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	output, err := svc.DescribeVpnGateways(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.VpnGateways
	return nil
}

func resolveVpnGatewayArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.VpnGateway)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "vpn-gateway/" + aws.ToString(item.VpnGatewayId),
	}
	return resource.Set(c.Name, a.String())
}
