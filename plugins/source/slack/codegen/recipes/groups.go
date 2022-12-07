package recipes

import (
	"github.com/slack-go/slack"
)

func GroupResources() []*Resource {
	resources := []*Resource{
		{
			DataStruct: &slack.GroupConversation{},
			PKColumns:  []string{"name"},
		},
	}
	for _, r := range resources {
		r.Service = "groups"
	}
	return resources
}
