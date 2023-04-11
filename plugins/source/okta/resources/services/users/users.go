package users

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:      "okta_users",
		Resolver:  fetchUsers,
		Transform: client.TransformWithStruct(&okta.User{}),
		Columns: []schema.Column{
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
