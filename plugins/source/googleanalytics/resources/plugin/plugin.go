package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleanalytics/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "googleanalytics"
	Kind    = "source"
	Team    = "cloudquery"
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"googleanalytics",
		Version,
		client.Configure,
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
	)
}
