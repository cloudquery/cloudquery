package users

import (
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func UserNotificationRules() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_user_notification_rules",
		Description: `https://developer.pagerduty.com/api-reference/043092de7e3e1-list-a-user-s-notification-rules`,
		Resolver:    fetchUserNotificationRules,
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "type",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "summary",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Summary"),
			},
			{
				Name:     "self",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Self"),
			},
			{
				Name:     "html_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "start_delay_in_minutes",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.PathResolver("StartDelayInMinutes"),
			},
			{
				Name:     "created_at",
				Type:     arrow.FixedWidthTypes.Timestamp_us,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "contact_method",
				Type:     types.ExtensionTypes.JSON,
				Resolver: schema.PathResolver("ContactMethod"),
			},
			{
				Name:     "urgency",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("Urgency"),
			},
		},
	}
}
