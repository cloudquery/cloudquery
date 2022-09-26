package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/resourcegroups/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"reflect"
	"strings"
)

func ResourceGroupsResources() []*Resource {
	resources := []*Resource{

		{
			SubService: "resource_groups",
			Struct:     &models.ResourceGroupWrapper{},
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
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
