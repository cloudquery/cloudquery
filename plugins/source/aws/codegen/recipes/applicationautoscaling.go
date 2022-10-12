package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApplicationAutoScalingResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "policies",
			Struct:      &types.ScalingPolicy{},
			Description: "https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingPolicy.html",
			SkipFields:  []string{"PolicyARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("PolicyARN")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "applicationautoscaling"
		r.Multiplex = `client.ServiceAccountRegionNamespaceMultiplexer("application-autoscaling")`
	}
	return resources
}
