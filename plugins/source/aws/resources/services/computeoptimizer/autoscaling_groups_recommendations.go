package computeoptimizer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func AutoscalingGroupsRecommendations() *schema.Table {
	tableName := "aws_computeoptimizer_autoscaling_group_recommendations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_AutoScalingGroupRecommendation.html`,
		Resolver:    fetchAutoscalingGroupsRecommendations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "compute-optimizer"),
		Transform:   transformers.TransformWithStruct(&types.AutoScalingGroupRecommendation{}, transformers.WithPrimaryKeys("AutoScalingGroupArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
		},
	}
}

func fetchAutoscalingGroupsRecommendations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	s := cl.Services()
	svc := s.Computeoptimizer

	input := computeoptimizer.GetAutoScalingGroupRecommendationsInput{
		MaxResults: aws.Int32(1000),
	}
	// No paginator available
	for {
		response, err := svc.GetAutoScalingGroupRecommendations(ctx, &input, func(options *computeoptimizer.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		if len(response.Errors) > 0 {
			cl.Logger().Error().Str("table", "aws_computeoptimizer_autoscaling_group_recommendations").Msgf("Errors in response: %v", response.Errors)
		}

		if response.AutoScalingGroupRecommendations != nil {
			res <- response.AutoScalingGroupRecommendations
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}

		input.NextToken = response.NextToken
	}

	return nil
}
