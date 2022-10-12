package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudwatchResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "alarms",
			Struct:      &types.MetricAlarm{},
			Description: "https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricAlarm.html",
			SkipFields:  []string{"AlarmArn", "Dimensions"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveCloudwatchAlarmTags`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("AlarmArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "dimensions",
						Type:     schema.TypeJSON,
						Resolver: `resolveCloudwatchAlarmDimensions`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "cloudwatch"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("logs")`
	}
	return resources
}
