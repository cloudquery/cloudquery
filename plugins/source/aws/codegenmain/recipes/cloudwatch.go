package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/plugin-sdk/codegen"
)

func init() {
	add(&Resource{
		DefaultColumns:  []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
		AWSStruct:       &types.MetricAlarm{},
		AWSService:      "Cloudwatch",
		Template:        "resource_get",
		ItemsStruct:     &cloudwatch.DescribeAlarmsOutput{},
		ColumnOverrides: map[string]codegen.ColumnDefinition{},
		// TODO query and add tags
	})
}
