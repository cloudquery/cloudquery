// Code generated by codegen; DO NOT EDIT.

package route53

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func HostedZoneQueryLoggingConfigs() *schema.Table {
	return &schema.Table{
		Name:      "aws_route53_hosted_zone_query_logging_configs",
		Resolver:  fetchRoute53HostedZoneQueryLoggingConfigs,
		Multiplex: client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveRoute53HostedZoneQueryLoggingConfigsArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "hosted_zone_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "cloud_watch_logs_log_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloudWatchLogsLogGroupArn"),
			},
			{
				Name:     "hosted_zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HostedZoneId"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
		},
	}
}
