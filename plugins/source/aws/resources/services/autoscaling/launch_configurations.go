package autoscaling

import (
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func LaunchConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "aws_autoscaling_launch_configurations",
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LaunchConfiguration.html`,
		Resolver:    fetchAutoscalingLaunchConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.LaunchConfiguration{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchConfigurationARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
