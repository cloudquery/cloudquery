package applications

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func ApplicationGroupAssignments() *schema.Table {
	return &schema.Table{
		Name:      "okta_application_group_assignments",
		Resolver:  fetchApplicationGroupAssignments,
		Transform: client.TransformWithStruct(&okta.ApplicationGroupAssignment{}),
		Columns: []schema.Column{
			{
				Name:     "app_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
