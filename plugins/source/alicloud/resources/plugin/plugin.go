package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/alicloud/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "alicloud"
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
