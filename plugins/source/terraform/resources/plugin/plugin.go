package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/cloudquery/plugins/source/terraform/resources/services"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"terraform",
		Version,
		[]*schema.Table{
			services.TFData(),
		},
		client.Configure,
	)
}
