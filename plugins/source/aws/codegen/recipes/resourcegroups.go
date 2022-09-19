package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/resourcegroups"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ResourceGroupsResources() []*Resource {
	resources := []*Resource{

		{
			SubService: "resource_groups",
			Struct:     &resourcegroups.ResourceGroupWrapper{},
			SkipFields: []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveResourcegroupsResourceGroupTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "resourcegroups"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("resource-groups")`
	}
	return resources
}
