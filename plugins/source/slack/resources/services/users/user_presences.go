package users

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
	"github.com/slack-go/slack"
)

func UserPresences() *schema.Table {
	return &schema.Table{
		Name:        "slack_user_presences",
		Description: `https://api.slack.com/methods/users.getPresence`,
		Resolver:    fetchUserPresences,
		Transform:   transformers.TransformWithStruct(&slack.UserPresence{}),
		Columns: []schema.Column{
			{
				Name:     "user_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "last_activity",
				Type:     schema.TypeTimestamp,
				Resolver: client.JSONTimeResolver("LastActivity"),
			},
		},
	}
}
