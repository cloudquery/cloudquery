package users

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/slack-go/slack"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "slack_users",
		Description: `https://api.slack.com/methods/users.list`,
		Resolver:    fetchUsers,
		Transform:   transformers.TransformWithStruct(&slack.User{}, transformers.WithSkipFields("Has2FA", "Presence")),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "updated",
				Type:     schema.TypeTimestamp,
				Resolver: client.JSONTimeResolver("Updated"),
			},
		},

		Relations: []*schema.Table{
			UserPresences(),
		},
	}
}
