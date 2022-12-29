package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var Version = "Development"

func Snyk() *source.Plugin {
	return source.NewPlugin(
		"snyk",
		Version,
		tables(),
		client.Configure,
	)
}
