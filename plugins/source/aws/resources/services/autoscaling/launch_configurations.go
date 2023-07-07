package autoscaling

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func LaunchConfigurations() *schema.Table {
	tableName := "aws_autoscaling_launch_configurations"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LaunchConfiguration.html`,
		Resolver:    fetchAutoscalingLaunchConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.LaunchConfiguration{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("LaunchConfigurationARN"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchAutoscalingLaunchConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling
	paginator := autoscaling.NewDescribeLaunchConfigurationsPaginator(svc, &autoscaling.DescribeLaunchConfigurationsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *autoscaling.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.LaunchConfigurations
	}
	return nil
}
