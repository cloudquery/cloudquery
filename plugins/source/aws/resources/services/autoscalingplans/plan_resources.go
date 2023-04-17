package autoscalingplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func planResources() *schema.Table {
	tableName := "aws_autoscaling_plan_resources"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingPlanResource.html`,
		Resolver:    fetchPlanResources,
		Transform:   transformers.TransformWithStruct(&types.ScalingPlanResource{}, transformers.WithPrimaryKeys("ScalingPlanName", "ResourceId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchPlanResources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Autoscalingplans
	p := parent.Item.(types.ScalingPlan)

	config := autoscalingplans.DescribeScalingPlanResourcesInput{
		ScalingPlanName: p.ScalingPlanName,
	}
	// No paginator available
	for {
		output, err := svc.DescribeScalingPlanResources(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.ScalingPlanResources

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
