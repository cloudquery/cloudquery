package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/mysql/client"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
)

var Version = "Development"

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"mysql",
		Version,
		nil,
		client.Configure,
		source.WithDynamicTableOption(getDynamicTables),
		source.WithNoInternalColumns(),
		source.WithUnmanaged(),
	)
}
