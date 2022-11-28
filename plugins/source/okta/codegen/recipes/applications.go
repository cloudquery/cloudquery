package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/groups/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

func Applications() []*Resource {
	return []*Resource{
		{
			DataStruct: &okta.Application{},
			PKColumns:  []string{"id"},
			Service:    "applications",
			Relations:  []string{"ApplicationUsers()", "ApplicationGroupAssignments()"},
		},
		{
			DataStruct: &models.ApplicationUser{},
			PKColumns:  []string{"app_id", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "app_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Service:               "applications",
			UnwrapEmbeddedStructs: true,
		},
		{
			DataStruct: &okta.ApplicationGroupAssignment{},
			PKColumns:  []string{"app_id", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "app_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Service: "applications",
		},
	}
}
