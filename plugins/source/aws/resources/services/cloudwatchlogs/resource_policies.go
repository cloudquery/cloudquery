// Code generated by codegen; DO NOT EDIT.

package cloudwatchlogs

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourcePolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudwatchlogs_resource_policies",
		Description: `https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ResourcePolicy.html`,
		Resolver:    fetchCloudwatchlogsResourcePolicies,
		Multiplex:   client.ServiceAccountRegionMultiplexer("logs"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "policy_document",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PolicyDocument"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyName"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
