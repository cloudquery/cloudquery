package savingsplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchSavingsPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	svc := meta.(*client.Client).Services().Savingsplans
	config := savingsplans.DescribeSavingsPlansInput{
		MaxResults: aws.Int32(1000),
	}
	for {
		response, err := svc.DescribeSavingsPlans(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.SavingsPlans
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
