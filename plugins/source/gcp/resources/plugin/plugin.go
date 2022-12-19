package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugins/source/gcp/client"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	allTables := PluginAutoGeneratedTables()
	// here you can append custom non-generated tables
	return source.NewPlugin(
		"gcp",
		Version,
		allTables,
		client.New,
	)
}
