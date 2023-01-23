package users

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Users() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_users",
		Description: `https://developer.pagerduty.com/api-reference/c96e889522dd6-list-users`,
		Resolver:    fetchUsers,
		Transform:   transformers.TransformWithStruct(&pagerduty.User{}, transformers.WithUnwrapAllEmbeddedStructs(), transformers.WithSkipFields("HTMLURL", "AvatarURL")),
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
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "avatar_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvatarURL"),
			},
		},

		Relations: []*schema.Table{
			UserContactMethods(),
			UserNotificationRules(),
		},
	}
}
