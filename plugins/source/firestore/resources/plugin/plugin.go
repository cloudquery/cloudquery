package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/firestore/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var version = "development"

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		"firestore",
		version,
		client.Configure,
	)
}
