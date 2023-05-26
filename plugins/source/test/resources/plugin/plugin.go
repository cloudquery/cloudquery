package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/test/client"
	"github.com/cloudquery/cloudquery/plugins/source/test/resources/services"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"test",
		Version,
		[]*schema.Table{
			services.TestSomeTable(),
			services.TestDataTable(),
		},
		client.New,
	)
}
