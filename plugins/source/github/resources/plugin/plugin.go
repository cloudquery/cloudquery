package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/github/client"
	"github.com/cloudquery/plugin-sdk/plugins"
)

const (
	exampleConfig = `
		access_token: <Access Token HERE>
		orgs:
		- cloudquery
`
)

var (
	Version = "Development"
)

func Plugin() *plugins.SourcePlugin {
	allTables := Tables()
	// here you can append custom non-generated tables
	return plugins.NewSourcePlugin(
		"github",
		Version,
		allTables,
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
		plugins.WithClassifyError(client.IgnoreError),
	)
}
