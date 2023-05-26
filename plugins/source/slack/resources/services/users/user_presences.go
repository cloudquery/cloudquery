package users

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/slack/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
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
				Name:       "user_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("id"),
				PrimaryKey: true,
			},
			{
				Name:     "last_activity",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: client.JSONTimeResolver("LastActivity"),
			},
		},
	}
}
