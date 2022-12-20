package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/groups/models"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func Groups() []*Resource {
	return []*Resource{
		{
			DataStruct: &okta.Group{},
			PKColumns:  []string{"id"},
			Service:    "groups",
			Relations:  []string{"GroupUsers()"},
		},
		{
			DataStruct: &models.GroupUser{},
			PKColumns:  []string{"group_id", "id"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "group_id",
					Type:     schema.TypeString,
					Resolver: "schema.ParentColumnResolver(\"id\")",
				},
			},
			Service: "groups",
		},
	}
}
