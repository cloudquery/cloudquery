package elbv2

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:        "aws_elbv2_load_balancers",
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_LoadBalancer.html`,
		Resolver:    fetchElbv2LoadBalancers,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
		Transform:   transformers.TransformWithStruct(&types.LoadBalancer{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "web_acl_arn",
				Type:     schema.TypeString,
				Resolver: resolveElbv2loadBalancerWebACLArn,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv2loadBalancerTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LoadBalancerArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			Listeners(),
			LoadBalancerAttributes(),
		},
	}
}
