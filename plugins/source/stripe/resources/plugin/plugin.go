package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/stripe/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"stripe",
		Version,
		tables(),
		client.Configure,
	)
}
