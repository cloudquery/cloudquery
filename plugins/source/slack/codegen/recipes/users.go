package recipes

import (
	"github.com/slack-go/slack"
)

func UserResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "users",
			Struct:      &slack.User{},
			Description: "", // TODO
			PKColumns:   []string{"id"},
		},
	}

	for _, r := range resources {
		r.Service = "users"
	}
	return resources
}
