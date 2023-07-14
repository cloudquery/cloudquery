package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/oracledb/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var Version = "Development"

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"oracledb",
		Version,
		client.Configure,
	)
}
