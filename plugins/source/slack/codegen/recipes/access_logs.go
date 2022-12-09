package recipes

import (
	"github.com/slack-go/slack"
)

func AccessLogResources() []*Resource {
	resources := []*Resource{
		{
			DataStruct: &slack.Login{},
			PKColumns:  []string{"user_id"},
		},
	}
	for _, r := range resources {
		r.Service = "access_logs"
		r.Multiplex = `client.TeamMultiplex`
		r.ImportClient = true
	}
	return resources
}
