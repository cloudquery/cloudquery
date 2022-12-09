package recipes

import (
	"github.com/cloudquery/cloudquery/plugins/source/slack/resources/services/conversations/models"
)

func ConversationResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "conversations",
			DataStruct: &models.Conversation{},
			PKColumns:  []string{"id"},
		},
	}
	for _, r := range resources {
		r.Service = "conversations"
		r.Multiplex = `client.TeamMultiplex`
	}
	return resources
}
