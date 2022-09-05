package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func init() {
	add(&Resource{
		DefaultColumns:             []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:                  &types.LogGroup{},
		AWSService:                 "CloudwatchLogs",
		MultiplexerServiceOverride: "logs",
		Template:                   "resource_get",
		ItemsStruct:                &cloudwatchlogs.DescribeLogGroupsOutput{},
		ColumnOverrides:            map[string]codegen.ColumnDefinition{},
		// TODO query and add tags
	},
		&Resource{
			DefaultColumns:             []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
			AWSStruct:                  &types.MetricFilter{},
			AWSService:                 "CloudwatchLogs",
			MultiplexerServiceOverride: "logs",
			Template:                   "resource_get",
			ItemsStruct:                &cloudwatchlogs.DescribeMetricFiltersOutput{},
			ColumnOverrides:            map[string]codegen.ColumnDefinition{},
		},
	)
}
