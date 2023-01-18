package autoscaling

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/autoscaling/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:        "aws_autoscaling_groups",
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AutoScalingGroup.html`,
		Resolver:    fetchAutoscalingGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("autoscaling"),
		Transform:   transformers.TransformWithStruct(&models.AutoScalingGroupWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
		},
		Relations: []*schema.Table{
			GroupScalingPolicies(),
			GroupLifecycleHooks(),
		},
	}
}
