package recipes

import (
	"github.com/slack-go/slack"
)

func ConversationResources() []*Resource {
	resources := []*Resource{
		{
			DataStruct: &slack.Conversation{},
			PKColumns:  []string{"id"},
		},
	}
	for _, r := range resources {
		r.Service = "conversations"
	}
	return resources
}
