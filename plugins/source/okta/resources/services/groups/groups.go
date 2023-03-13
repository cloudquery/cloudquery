package groups

import (
	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "okta_groups",
		Resolver:  fetchGroups,
		Transform: client.TransformWithStruct(&okta.Group{}),
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

		Relations: []*schema.Table{
			GroupUsers(),
		},
	}
}
