package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"salesforce",
		Version,
		Configure,
	)
}
