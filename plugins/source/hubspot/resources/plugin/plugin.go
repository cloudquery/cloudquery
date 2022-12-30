package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/hubspot/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var Version = "Development"

func HubSpot() *source.Plugin {
	return source.NewPlugin(
		"hubspot",
		Version,
		tables(),
		client.Configure,
	)
}
