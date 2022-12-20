package recipes

import "github.com/PagerDuty/go-pagerduty"

func UserResources() []*Resource {
	return []*Resource{
		{
			SubService:  "users",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.User{},
			Description: "https://developer.pagerduty.com/api-reference/c96e889522dd6-list-users",

			Relations: []string{"UserContactMethods()", "UserNotificationRules()"},
		},
		{
			SubService:  "user_contact_methods",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.ContactMethod{},
			Description: "https://developer.pagerduty.com/api-reference/50d46c0eb020d-list-a-user-s-contact-methods",

			Template:                 "nested_no_options",
			ListFunctionNameOverride: "ListUserContactMethodsWithContext",
			ResponseFieldOverride:    "ContactMethods",
		},
		{
			SubService:  "user_notification_rules",
			PKColumns:   []string{"id"},
			Struct:      pagerduty.NotificationRule{},
			Description: "https://developer.pagerduty.com/api-reference/043092de7e3e1-list-a-user-s-notification-rules",

			Template:                 "nested_no_options",
			ListFunctionNameOverride: "ListUserNotificationRulesWithContext",
			ResponseFieldOverride:    "NotificationRules",
			ResponseStructOverride:   "ListUserNotificationRulesResponse",
		},
	}
}
