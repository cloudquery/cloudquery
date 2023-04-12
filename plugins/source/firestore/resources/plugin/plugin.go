package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/firestore/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var version = "v1"

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"firestore",
		version,
		nil,
		client.Configure,
		source.WithDynamicTableOption(getDynamicTables),
		source.WithNoInternalColumns(),
		source.WithUnmanaged(),
	)
}
