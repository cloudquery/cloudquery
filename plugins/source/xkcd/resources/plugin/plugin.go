package plugin

import (
	internalPlugin "github.com/cloudquery/cloudquery/plugins/source/xkcd/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		internalPlugin.Name,
		internalPlugin.Version,
		Configure,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
	)
}
