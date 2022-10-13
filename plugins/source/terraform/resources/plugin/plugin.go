package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/cloudquery/plugins/source/terraform/resources/services"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"terraform",
		Version,
		[]*schema.Table{
			services.TFData(),
		},
		client.Configure,
	)
}
