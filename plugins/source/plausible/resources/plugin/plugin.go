package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/plausible/client"
	"github.com/cloudquery/cloudquery/plugins/source/plausible/resources/services"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "Development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"plausible",
		Version,
		[]*schema.Table{
			services.StatsTimeseries(),
		},
		client.Configure,
	)
}
