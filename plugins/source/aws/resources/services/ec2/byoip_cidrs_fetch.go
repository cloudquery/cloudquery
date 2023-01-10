package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEc2ByoipCidrs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	config := ec2.DescribeByoipCidrsInput{
		MaxResults: aws.Int32(100),
	}

	c := meta.(*client.Client)
	// DescribeByoipCidrs does not work in next regions, so we ignore them.
	if _, ok := map[string]struct{}{
		"cn-north-1":     {},
		"cn-northwest-1": {},
	}[c.Region]; ok {
		return nil
	}
	svc := c.Services().Ec2
	for {
		response, err := svc.DescribeByoipCidrs(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.ByoipCidrs
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
