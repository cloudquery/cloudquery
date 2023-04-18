package applicationautoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func ScalingActivities() *schema.Table {
	tableName := "aws_applicationautoscaling_scaling_activities"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/application/APIReference/API_ScalingActivity.html`,
		Resolver:    fetchScalingActivities,
		Multiplex:   client.ServiceAccountRegionNamespaceMultiplexer(tableName, "application-autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.ScalingActivity{}, transformers.WithPrimaryKeys("ResourceId")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchScalingActivities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Applicationautoscaling

	config := applicationautoscaling.DescribeScalingActivitiesInput{
		ServiceNamespace: types.ServiceNamespace(c.AutoscalingNamespace),
	}
	paginator := applicationautoscaling.NewDescribeScalingActivitiesPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.ScalingActivities
	}
	return nil
}
