package users

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func UserNotificationRules() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_user_notification_rules",
		Description: `https://developer.pagerduty.com/api-reference/043092de7e3e1-list-a-user-s-notification-rules`,
		Resolver:    fetchUserNotificationRules,
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
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "summary",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Summary"),
			},
			{
				Name:     "self",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Self"),
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "start_delay_in_minutes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("StartDelayInMinutes"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "contact_method",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ContactMethod"),
			},
			{
				Name:     "urgency",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Urgency"),
			},
		},
	}
}
