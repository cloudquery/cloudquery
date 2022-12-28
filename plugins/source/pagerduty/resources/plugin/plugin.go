package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/pagerduty/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var (
	Version = "Development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"pagerduty",
		Version,
		AllTables(),
		client.Configure,
	)
}
