package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/mysql/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var Version = "Development"

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"mysql",
		Version,
		client.Configure,
	)
}
