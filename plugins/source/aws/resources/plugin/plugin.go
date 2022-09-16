package plugin

import (
	_ "embed"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/plugins"
)

var (
	Version = "Development"

	//go:embed example.yml
	exampleConfig string
)

func AWS() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"aws",
		Version,
		tables(),
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
	)
}
