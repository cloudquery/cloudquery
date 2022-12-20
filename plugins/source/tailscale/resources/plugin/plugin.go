package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var Version = "Development"

func Tailscale() *source.Plugin {
	return source.NewPlugin(
		"tailscale",
		Version,
		tables(),
		client.Configure,
	)
}
