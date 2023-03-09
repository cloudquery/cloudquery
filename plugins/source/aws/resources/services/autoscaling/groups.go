package autoscaling

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/autoscaling/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Groups() *schema.Table {
	tableName := "aws_autoscaling_groups"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AutoScalingGroup.html`,
		Resolver:    fetchAutoscalingGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "autoscaling"),
		Transform:   transformers.TransformWithStruct(&models.AutoScalingGroupWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "load_balancers",
				Type:     schema.TypeJSON,
				Resolver: resolveAutoscalingGroupLoadBalancers,
			},
			{
				Name:     "load_balancer_target_groups",
				Type:     schema.TypeJSON,
				Resolver: resolveAutoscalingGroupLoadBalancerTargetGroups,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "tags_raw",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
		Relations: []*schema.Table{
			GroupScalingPolicies(),
			GroupLifecycleHooks(),
		},
	}
}
