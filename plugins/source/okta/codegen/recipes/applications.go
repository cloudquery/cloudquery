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
			SkipFields: []string{"Profile"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "profile",
					Type:     schema.TypeJSON,
					Resolver: `schema.PathResolver("Profile")`,
				},
			},
			Relations: []string{"ApplicationUsers()", "ApplicationGroupAssignments()"},
		},
		{
			DataStruct: &models.ApplicationUser{},
			PKColumns:  []string{"app_id", "id"},
			SkipFields: []string{"Profile"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "app_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("id")`,
				},
				{
					Name:     "profile",
					Type:     schema.TypeJSON,
					Resolver: `schema.PathResolver("Profile")`,
				},
			},
			Service:               "applications",
			UnwrapEmbeddedStructs: true,
		},
		{
			DataStruct: &okta.ApplicationGroupAssignment{},
			PKColumns:  []string{"app_id", "id"},
			SkipFields: []string{"Profile"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "app_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("id")`,
				},
				{
					Name:     "profile",
					Type:     schema.TypeJSON,
					Resolver: `schema.PathResolver("Profile")`,
				},
			},
			Service: "applications",
		},
	}
}
