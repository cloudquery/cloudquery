package applicationautoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ScalableTargets() *schema.Table {
	tableName := "aws_applicationautoscaling_scalable_targets"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalableTarget.html`,
		Resolver:    fetchScalableTargets,
		Multiplex:   client.ServiceAccountRegionNamespaceMultiplexer(tableName, "application-autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.ScalableTarget{}, transformers.WithPrimaryKeys("ResourceId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchScalableTargets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Applicationautoscaling

	config := applicationautoscaling.DescribeScalableTargetsInput{
		ServiceNamespace: types.ServiceNamespace(c.AutoscalingNamespace),
	}
	for {
		output, err := svc.DescribeScalableTargets(ctx, &config)
		if err != nil {
			return err
		}

		res <- output.ScalableTargets

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
