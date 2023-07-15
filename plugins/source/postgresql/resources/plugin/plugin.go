package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/postgresql/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var Version = "Development"

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"postgresql",
		Version,
		client.Configure,
	)
}
