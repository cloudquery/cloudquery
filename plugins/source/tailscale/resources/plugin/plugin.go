package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/plugins"
)

var Version = "Development"

func Tailscale() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"tailscale",
		Version,
		tables(),
		client.Configure,
	)
}
