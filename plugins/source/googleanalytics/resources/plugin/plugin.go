package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/googleanalytics/client"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
)

var Version = "Development"

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"googleanalytics",
		Version,
		nil,
		client.Configure,
		source.WithDynamicTableOption(client.GetTables),
	)
}
