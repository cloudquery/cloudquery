package users

import (
	"github.com/cloudquery/cloudquery/plugins/source/gitlab/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/xanzy/go-gitlab"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "gitlab_users",
		Resolver:  fetchUsers,
		Transform: client.TransformWithStruct(&gitlab.User{}),
		Columns: []schema.Column{
			client.BaseURLColumn,
			{
				Name:          "last_activity_on",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("LastActivityOn"),
				IgnoreInTests: true,
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
