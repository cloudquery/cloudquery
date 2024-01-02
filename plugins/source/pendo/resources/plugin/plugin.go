package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "pendo"
	Kind    = "source"
	Team    = "cloudquery"
	Version = "v0.0.1"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(Name, Version, Configure, plugin.WithKind(Kind), plugin.WithTeam(Team))
}
