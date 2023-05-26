package users

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
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
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "updated",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: client.JSONTimeResolver("Updated"),
			},
		},

		Relations: []*schema.Table{
			UserPresences(),
		},
	}
}
