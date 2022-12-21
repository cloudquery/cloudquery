// Code generated by codegen; DO NOT EDIT.

package groups

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func GroupUsers() *schema.Table {
	return &schema.Table{
		Name:     "okta_group_users",
		Resolver: fetchGroupUsers,
		Columns: []schema.Column{
			{
				Name:     "group_id",
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
