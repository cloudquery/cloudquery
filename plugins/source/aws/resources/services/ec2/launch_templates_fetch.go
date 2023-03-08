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

func fetchEc2LaunchTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config ec2.DescribeLaunchTemplatesInput
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	for {
		output, err := svc.DescribeLaunchTemplates(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.LaunchTemplates
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}

func resolveEc2LaunchTemplateArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.LaunchTemplate)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "launch-template/" + aws.ToString(item.LaunchTemplateId),
	}
	return resource.Set(c.Name, a.String())
}
