package plugin

import (
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var Version = "Development"

func Snyk() *plugin.Plugin {
	return plugin.NewPlugin(
		"snyk",
		Version,
		configure,
	)
}
