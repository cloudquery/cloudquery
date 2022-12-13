package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/client"
	"github.com/cloudquery/plugin-sdk/plugins"
)

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"azuredevops",
		Version,
		tables(),
		client.New,
	)
}
