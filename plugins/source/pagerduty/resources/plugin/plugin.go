package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "pagerduty"
	Kind    = "source"
	Team    = "cloudquery"
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		Name,
		Version,
		Configure,
		plugin.WithJSONSchema(client.JSONSchema),
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
	)
}
