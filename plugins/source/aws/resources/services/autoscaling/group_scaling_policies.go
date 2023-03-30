package autoscaling

import (
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func GroupScalingPolicies() *schema.Table {
	tableName := "aws_autoscaling_group_scaling_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_ScalingPolicy.html`,
		Resolver:    fetchAutoscalingGroupScalingPolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "autoscaling"),
		Transform:   transformers.TransformWithStruct(&types.ScalingPolicy{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "group_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
