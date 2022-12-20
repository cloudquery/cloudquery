package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/plugins"
)

var (
	Version = "Development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"pagerduty",
		Version,
		AllTables(),
		client.Configure,
	)
}
