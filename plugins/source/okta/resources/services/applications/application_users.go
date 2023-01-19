package applications

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func ApplicationUsers() *schema.Table {
	return &schema.Table{
		Name:      "okta_application_users",
		Resolver:  fetchApplicationUsers,
		Transform: transformers.TransformWithStruct(&okta.AppUser{}, client.SharedTransformers()...),
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
