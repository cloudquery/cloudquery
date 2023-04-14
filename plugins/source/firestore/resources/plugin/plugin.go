package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/firestore/client"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
)

var version = "development"

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
