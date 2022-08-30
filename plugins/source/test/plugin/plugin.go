package plugin

import (
	"github.com/cloudquery/cq-source-test/client"
	"github.com/cloudquery/cq-source-test/tables"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

const exampleConfig = `
# This is an example config file for the test plugin.
account_ids: []
`

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"test",
		Version,
		[]*schema.Table{
			tables.TestSomeTable(),
		},
		client.Configure,
		plugins.WithSourceExampleConfig(exampleConfig),
	)
}
