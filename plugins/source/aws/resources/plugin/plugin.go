package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var (
	Version = "Development"
)

func AWS() *source.Plugin {
	return source.NewPlugin(
		"aws",
		Version,
		tables(),
		client.Configure,
	)
}
