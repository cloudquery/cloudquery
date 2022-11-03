package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSsmParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm
	params := ssm.DescribeParametersInput{}
	for {
		output, err := svc.DescribeParameters(ctx, &params)
		if err != nil {
			return err
		}
		res <- output.Parameters
		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}
