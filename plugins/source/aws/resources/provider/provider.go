package provider

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/plugins"
)

var (

	Version = "Development"
)

func Provider() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"aws",
		Version,
		tables(),
		client.Configure,
	)
}
