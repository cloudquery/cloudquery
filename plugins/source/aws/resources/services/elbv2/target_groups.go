// Code generated by codegen; DO NOT EDIT.

package elbv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TargetGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_elbv2_target_groups",
		Description: `https://docs.aws.amazon.com/elasticloadbalancing/latest/APIReference/API_TargetGroup.html`,
		Resolver:    fetchElbv2TargetGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
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
			{
				Name:     "health_check_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HealthCheckEnabled"),
			},
			{
				Name:     "health_check_interval_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheckIntervalSeconds"),
			},
			{
				Name:     "health_check_path",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckPath"),
			},
			{
				Name:     "health_check_port",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckPort"),
			},
			{
				Name:     "health_check_protocol",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckProtocol"),
			},
			{
				Name:     "health_check_timeout_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheckTimeoutSeconds"),
			},
			{
				Name:     "healthy_threshold_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthyThresholdCount"),
			},
			{
				Name:     "ip_address_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpAddressType"),
			},
			{
				Name:     "load_balancer_arns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("LoadBalancerArns"),
			},
			{
				Name:     "matcher",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Matcher"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
			},
			{
				Name:     "protocol",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Protocol"),
			},
			{
				Name:     "protocol_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProtocolVersion"),
			},
			{
				Name:     "target_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetGroupName"),
			},
			{
				Name:     "target_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TargetType"),
			},
			{
				Name:     "unhealthy_threshold_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("UnhealthyThresholdCount"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},

		Relations: []*schema.Table{
			TargetGroupTargetHealthDescriptions(),
		},
	}
}
