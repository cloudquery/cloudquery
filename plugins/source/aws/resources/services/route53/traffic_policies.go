// Code generated by codegen; DO NOT EDIT.

package route53

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TrafficPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_route53_traffic_policies",
		Description: "https://docs.aws.amazon.com/Route53/latest/APIReference/API_TrafficPolicySummary.html",
		Resolver:    fetchRoute53TrafficPolicies,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveTrafficPolicyArn(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "latest_version",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("LatestVersion"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "traffic_policy_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TrafficPolicyCount"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},

		Relations: []*schema.Table{
			TrafficPolicyVersions(),
		},
	}
}
