package groups

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func Groups() *schema.Table {
	return &schema.Table{
		Name:      "okta_groups",
		Resolver:  fetchGroups,
		Transform: transformers.TransformWithStruct(&okta.Group{}),
		Columns: []schema.Column{
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
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

		Relations: []*schema.Table{
			GroupUsers(),
		},
	}
}
