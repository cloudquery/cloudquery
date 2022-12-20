package recipes

import (
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/slack-go/slack"
)

func FileResources() []*Resource {
	resources := []*Resource{
		{
			DataStruct:   &slack.File{},
			Description:  "https://api.slack.com/methods/files.list",
			PKColumns:    []string{"team_id", "id"},
			ExtraColumns: []codegen.ColumnDefinition{TeamIDColumn},
		},
	}
	for _, r := range resources {
		r.Service = "files"
		r.Multiplex = `client.TeamMultiplex`
	}
	return resources
}
