package recipes

import (
	"github.com/fastly/go-fastly/v7/fastly"
)

func AccountResources() []*Resource {
	resources := []*Resource{
		{
			TableName:   "account_users",
			DataStruct:  &fastly.User{},
			Description: "https://developer.fastly.com/reference/api/account/user/",
			PKColumns:   []string{"id"},
		},
		{
			TableName:   "account_events",
			DataStruct:  &fastly.Event{},
			Description: "https://developer.fastly.com/reference/api/account/events/",
			PKColumns:   []string{"id"},
		},
	}
	for _, r := range resources {
		r.Service = "account"
	}
	return resources
}
