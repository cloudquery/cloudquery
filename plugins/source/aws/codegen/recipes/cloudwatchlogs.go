package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudWatchLogsResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "resource_policies",
			Struct:      &types.ResourcePolicy{},
			Description: "https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_ResourcePolicy.html",
			PKColumns:   []string{"account_id", "region", "policy_name"},
			SkipFields:  []string{"PolicyDocument"},
			ExtraColumns: append(defaultRegionalColumns, []codegen.ColumnDefinition{
				{
					Name:     "policy_document",
					Type:     schema.TypeJSON,
					Resolver: `schema.PathResolver("PolicyDocument")`,
				},
			}...),
		},
		{
			SubService:  "metric_filters",
			Struct:      &types.MetricFilter{},
			Description: "https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_MetricFilter.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveMetricFilterArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "log_groups",
			Struct:      &types.LogGroup{},
			Description: "https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_LogGroup.html",
			SkipFields:  []string{"Arn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Arn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveLogGroupTags`,
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "cloudwatchlogs"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("logs")`
	}
	return resources
}
