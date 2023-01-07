package users

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/slack-go/slack"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "slack_users",
		Description: `https://api.slack.com/methods/users.list`,
		Resolver:    fetchUsers,
		Transform:   transformers.TransformWithStruct(&slack.User{}),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			UserPresences(),
		},
	}
}
