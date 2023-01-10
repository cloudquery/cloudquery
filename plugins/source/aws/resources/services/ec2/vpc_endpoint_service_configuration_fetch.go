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

func fetchEc2VpcEndpointServiceConfigurations(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	var config ec2.DescribeVpcEndpointServiceConfigurationsInput
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	for {
		output, err := svc.DescribeVpcEndpointServiceConfigurations(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.ServiceConfigurations
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveVpcEndpointServiceConfigurationArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.ServiceConfiguration)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "vpc_endpoint_service_configuration/" + aws.ToString(item.ServiceId),
	}
	return resource.Set(c.Name, a.String())
}
