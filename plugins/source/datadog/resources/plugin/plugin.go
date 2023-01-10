package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/datadog/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var (
	Version = "Development"
)

func Plugin() *source.Plugin {
	allTables := Tables()
	// here you can append custom non-generated tables
	return source.NewPlugin(
		"datadog",
		Version,
		allTables,
		client.Configure,
	)
}
