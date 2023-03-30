package autoscaling

import (
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GroupLifecycleHooks() *schema.Table {
	tableName := "aws_autoscaling_group_lifecycle_hooks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LifecycleHook.html`,
		Resolver:    fetchAutoscalingGroupLifecycleHooks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "autoscaling"),
		Transform:   client.TransformWithStruct(&types.LifecycleHook{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},
	}
}
