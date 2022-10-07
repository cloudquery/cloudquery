// Code generated by codegen; DO NOT EDIT.

package cloudtrail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Trails() *schema.Table {
	return &schema.Table{
		Name:      "aws_cloudtrail_trails",
		Resolver:  fetchCloudtrailTrails,
		Multiplex: client.ServiceAccountRegionMultiplexer("cloudtrail"),
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
				Name:     "cloudwatch_logs_log_group_name",
				Type:     schema.TypeString,
				Resolver: resolveCloudtrailTrailCloudwatchLogsLogGroupName,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TrailARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "status",
				Type:     schema.TypeJSON,
				Resolver: resolveCloudTrailStatus,
			},
			{
				Name:     "cloud_watch_logs_log_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloudWatchLogsLogGroupArn"),
			},
			{
				Name:     "cloud_watch_logs_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CloudWatchLogsRoleArn"),
			},
			{
				Name:     "has_custom_event_selectors",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HasCustomEventSelectors"),
			},
			{
				Name:     "has_insight_selectors",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HasInsightSelectors"),
			},
			{
				Name:     "home_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HomeRegion"),
			},
			{
				Name:     "include_global_service_events",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IncludeGlobalServiceEvents"),
			},
			{
				Name:     "is_multi_region_trail",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsMultiRegionTrail"),
			},
			{
				Name:     "is_organization_trail",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsOrganizationTrail"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "log_file_validation_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("LogFileValidationEnabled"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "s3_bucket_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("S3BucketName"),
			},
			{
				Name:     "s3_key_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("S3KeyPrefix"),
			},
			{
				Name:     "sns_topic_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnsTopicARN"),
			},
			{
				Name:     "sns_topic_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnsTopicName"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},

		Relations: []*schema.Table{
			TrailEventSelectors(),
		},
	}
}
