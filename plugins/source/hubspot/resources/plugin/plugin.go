package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "hubspot"
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
		plugin.WithJSONSchema(spec.JSONSchema),
	)
}
