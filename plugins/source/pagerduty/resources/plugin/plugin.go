package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Version = "Development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"pagerduty",
		Version,
		Configure,
	)
}
