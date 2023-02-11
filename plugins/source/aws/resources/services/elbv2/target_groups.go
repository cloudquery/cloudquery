package elbv2

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TargetGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_elbv2_target_groups",
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroup.html`,
		Resolver:    fetchElbv2TargetGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&types.TargetGroup{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv2targetGroupTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetGroupArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			TargetGroupTargetHealthDescriptions(),
		},
	}
}
