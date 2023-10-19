package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/postgresql/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "postgresql"
	Kind    = "source"
	Team    = "cloudquery"
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		Name,
		Version,
		client.Configure,
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
	)
}
