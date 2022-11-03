package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchEc2Eips(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	output, err := svc.DescribeAddresses(ctx, &ec2.DescribeAddressesInput{
		Filters: []types.Filter{{Name: aws.String("domain"), Values: []string{"vpc"}}},
	})
	if err != nil {
		return err
	}
	res <- output.Addresses
	return nil
}
