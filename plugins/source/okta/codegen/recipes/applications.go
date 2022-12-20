package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
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
			DataStruct: &okta.AppUser{},
			PKColumns:  []string{"app_id", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "app_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("id")`,
				},
			},
			Service:               "applications",
			UnwrapEmbeddedStructs: true,
			TableName:             "application_users",
		},
		{
			DataStruct: &okta.ApplicationGroupAssignment{},
			PKColumns:  []string{"app_id", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "app_id",
					Type:     schema.TypeString,
					Resolver: `schema.ParentColumnResolver("id")`,
				},
			},
			Service: "applications",
		},
	}
}
