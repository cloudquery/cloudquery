package elbv1

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LoadBalancerPolicies() *schema.Table {
	tableName := "aws_elbv1_load_balancer_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/2012-06-01/APIReference/API_PolicyDescription.html`,
		Resolver:    fetchElbv1LoadBalancerPolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "elasticloadbalancing"),
		Transform:   client.TransformWithStruct(&types.PolicyDescription{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "load_balancer_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "load_balancer_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("load_balancer_name"),
			},
			{
				Name:     "policy_attribute_descriptions",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv1loadBalancerPolicyAttributeDescriptions,
			},
		},
	}
}
