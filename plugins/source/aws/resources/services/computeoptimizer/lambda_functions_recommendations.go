package computeoptimizer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func LambdaFunctionsRecommendations() *schema.Table {
	tableName := "aws_computeoptimizer_lambda_function_recommendations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_LambdaFunctionRecommendation.html`,
		Resolver:    fetchLambdaFunctionsRecommendations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "compute-optimizer"),
		Transform:   transformers.TransformWithStruct(&types.LambdaFunctionRecommendation{}, transformers.WithPrimaryKeys("FunctionArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
		},
	}
}

func fetchLambdaFunctionsRecommendations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	s := c.Services()
	svc := s.Computeoptimizer

	input := computeoptimizer.GetLambdaFunctionRecommendationsInput{
		MaxResults: aws.Int32(1000),
	}
	// No paginator available
	for {
		response, err := svc.GetLambdaFunctionRecommendations(ctx, &input)
		if err != nil {
			return err
		}

		if response.LambdaFunctionRecommendations != nil {
			res <- response.LambdaFunctionRecommendations
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}

		input.NextToken = response.NextToken
	}

	return nil
}
