package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchFrauddetectorModelVersions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().FraudDetector

	input := &frauddetector.DescribeModelVersionsInput{ModelId: parent.Item.(*types.Model).ModelId}
	for {
		output, err := svc.DescribeModelVersions(ctx, input)
		if err != nil {
			return err
		}

		res <- output.ModelVersionDetails

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}

	return nil
}
