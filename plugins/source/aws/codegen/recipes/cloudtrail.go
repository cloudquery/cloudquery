package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudtrailResources() []*Resource {
	resources := []*Resource{
		{
			SubService:           "trails",
			Struct:               &types.Trail{},
			SkipFields:           []string{"TrailARN"},
			PostResourceResolver: `postCloudtrailTrailResolver`,
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
				},
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
			},
			Relations: []string{
				"TrailEventSelectors()",
			},
		},
		{
			SubService: "trail_event_selectors",
			Struct:     &types.EventSelector{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultAccountColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "trail_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "cloudtrail"
		r.Multiplex = `client.AccountMultiplex`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
