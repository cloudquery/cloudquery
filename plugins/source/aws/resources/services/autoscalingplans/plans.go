package autoscalingplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Plans() *schema.Table {
	tableName := "aws_autoscaling_plans"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingPlan.html`,
		Resolver:    fetchPlans,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "autoscaling-plans"),
		Transform:   transformers.TransformWithStruct(&types.ScalingPlan{}, transformers.WithPrimaryKeys("ScalingPlanName")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
		Relations: []*schema.Table{
			planResources(),
		},
	}
}

func fetchPlans(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscalingplans
	config := autoscalingplans.DescribeScalingPlansInput{}
	// No paginator available
	for {
		output, err := svc.DescribeScalingPlans(ctx, &config, func(options *autoscalingplans.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}

		res <- output.ScalingPlans

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
