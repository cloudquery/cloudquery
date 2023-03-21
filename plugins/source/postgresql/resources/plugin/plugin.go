package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/postgresql/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var Version = "Development"

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"postgresql",
		Version,
		nil,
		client.Configure,
		source.WithDynamicTableOption(getDynamicTables),
		source.WithNoInternalColumns(),
		source.WithUnmanaged(),
	)
}
