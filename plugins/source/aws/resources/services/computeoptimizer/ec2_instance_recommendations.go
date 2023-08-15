package computeoptimizer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func Ec2InstanceRecommendations() *schema.Table {
	tableName := "aws_computeoptimizer_ec2_instance_recommendations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/compute-optimizer/latest/APIReference/API_InstanceRecommendation.html`,
		Resolver:    fetchEc2InstanceRecommendations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "compute-optimizer"),
		Transform:   transformers.TransformWithStruct(&types.InstanceRecommendation{}, transformers.WithPrimaryKeys("InstanceArn")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "tags",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchEc2InstanceRecommendations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Computeoptimizer

	input := computeoptimizer.GetEC2InstanceRecommendationsInput{
		MaxResults: aws.Int32(1000),
	}
	// No paginator available
	for {
		response, err := svc.GetEC2InstanceRecommendations(ctx, &input, func(options *computeoptimizer.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		if len(response.Errors) > 0 {
			cl.Logger().Error().Str("table", "aws_computeoptimizer_ec2_instance_recommendations").Msgf("Errors in response: %v", response.Errors)
		}

		if response.InstanceRecommendations != nil {
			res <- response.InstanceRecommendations
		}

		if aws.ToString(response.NextToken) == "" {
			break
		}

		input.NextToken = response.NextToken
	}

	return nil
}
