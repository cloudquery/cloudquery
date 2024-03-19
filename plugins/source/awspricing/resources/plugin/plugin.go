package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/awspricing/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "awspricing"
	Kind    = "source"
	Team    = "cloudquery"
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		Name,
		Version,
		Configure,
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
		plugin.WithJSONSchema(client.JSONSchema),
	)
}
