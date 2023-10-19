package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "snyk"
	Kind    = "source"
	Team    = "cloudquery"
	Version = "development"
)

func Snyk() *plugin.Plugin {
	return plugin.NewPlugin(
		Name,
		Version,
		configure,
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
	)
}
