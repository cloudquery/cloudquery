package users

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_users",
		Description: `https://developer.pagerduty.com/api-reference/c96e889522dd6-list-users`,
		Resolver:    fetchUsers,
		Transform:   transformers.TransformWithStruct(&pagerduty.User{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL", "AvatarURL")),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "html_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "avatar_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("AvatarURL"),
			},
		},

		Relations: []*schema.Table{
			UserContactMethods(),
			UserNotificationRules(),
		},
	}
}
