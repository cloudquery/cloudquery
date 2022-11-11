package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/cloudtrail/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudtrailResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "trails",
			Struct:      &models.CloudTrailWrapper{},
			Description: "https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_Trail.html",
			SkipFields:  []string{"TrailARN"},
			ExtraColumns: append(defaultRegionalColumns, []codegen.ColumnDefinition{
				{
					Name:     "cloudwatch_logs_log_group_name",
					Type:     schema.TypeString,
					Resolver: `resolveCloudtrailTrailCloudwatchLogsLogGroupName`,
				},
				{
					Name:     "arn",
					Type:     schema.TypeString,
					Resolver: `schema.PathResolver("TrailARN")`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "status",
					Type:     schema.TypeJSON,
					Resolver: `resolveCloudTrailStatus`,
				},
			}...),
			Relations: []string{
				"TrailEventSelectors()",
			},
		},
		{
			SubService:  "trail_event_selectors",
			Struct:      &types.EventSelector{},
			Description: "https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_EventSelector.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "trail_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "cloudtrail"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("cloudtrail")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
