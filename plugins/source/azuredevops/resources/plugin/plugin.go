package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/azuredevops/client"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"azuredevops",
		Version,
		tables(),
		client.New,
	)
}
