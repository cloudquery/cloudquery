package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/fastly/client"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"fastly",
		Version,
		tables(),
		client.Configure,
	)
}
