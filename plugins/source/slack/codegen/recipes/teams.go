package recipes

import (
	"github.com/slack-go/slack"
)

func TeamResources() []*Resource {
	resources := []*Resource{
		{
			TableName:   "teams",
			DataStruct:  &slack.TeamInfo{},
			Description: "https://slack.com/api/team.info",
			PKColumns:   []string{"id"},
		},
	}
	for _, r := range resources {
		r.Service = "teams"
		r.Multiplex = `client.TeamMultiplex`
	}
	return resources
}
