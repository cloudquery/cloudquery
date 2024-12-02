package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/v3/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "hackernews"
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
