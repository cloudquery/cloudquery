package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "{{.Name}}"
	Kind    = "source"
	Team    = "{{.Org}}"
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(Name, Version, Configure, plugin.WithKind(Kind), plugin.WithTeam(Team))
}
