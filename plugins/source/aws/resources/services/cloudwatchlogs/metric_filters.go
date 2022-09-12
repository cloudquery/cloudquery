// Code generated by codegen; DO NOT EDIT.

package cloudwatchlogs

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func MetricFilters() *schema.Table {
	return &schema.Table{
		Name:      "aws_cloudwatchlogs_metric_filters",
		Resolver:  fetchCloudwatchlogsMetricFilters,
		Multiplex: client.ServiceAccountRegionMultiplexer("cloudwatchlogs"),
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
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveMetricFilterArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "filter_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FilterName"),
			},
			{
				Name:     "filter_pattern",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FilterPattern"),
			},
			{
				Name:     "log_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogGroupName"),
			},
			{
				Name:     "metric_transformations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MetricTransformations"),
			},
		},
	}
}
