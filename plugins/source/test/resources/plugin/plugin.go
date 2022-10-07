package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/test/client"
	"github.com/cloudquery/cloudquery/plugins/source/test/tables"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

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
		client.New,
	)
}
