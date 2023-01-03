package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"hackernews",
		Version,
		tables(),
		client.Configure,
	)
}
