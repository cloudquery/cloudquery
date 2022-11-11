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

func fetchEc2KeyPairs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeKeyPairsInput
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	output, err := svc.DescribeKeyPairs(ctx, &config)
	if err != nil {
		return err
	}
	res <- output.KeyPairs
	return nil
}

func resolveKeyPairArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	item := resource.Item.(types.KeyPairInfo)
	a := arn.ARN{
		Partition: cl.Partition,
		Service:   "ec2",
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "key-pair/" + aws.ToString(item.KeyPairId),
	}
	return resource.Set(c.Name, a.String())
}
