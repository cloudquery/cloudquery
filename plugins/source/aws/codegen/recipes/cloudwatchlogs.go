package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)



func CloudWatchLogsResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "metric_filters",
			Struct: &types.MetricFilter{},
			Multiplex: `client.ServiceAccountRegionMultiplexer("cloudwatchlogs")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `resolveMetricFilterArn`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
		},
		{
			SubService: "log_groups",
			Struct: &types.LogGroup{},
			Multiplex: `client.ServiceAccountRegionMultiplexer("cloudwatchlogs")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
				{
					Name: "arn",
					Type: schema.TypeString,
					Resolver: `schema.PathResolver("Arn")`,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name: "tags",
					Type: schema.TypeJSON,
					Resolver: `resolveLogGroupArn`,
				},
			}...),
		},
	}

	for _, r := range resources {
		r.Service = "cloudwatchlogs"
	}
	return resources
}