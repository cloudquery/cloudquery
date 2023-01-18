package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/resources/services/cohorts"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/resources/services/funnels"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"mixpanel",
		Version,
		[]*schema.Table{
			cohorts.Cohorts(),
			funnels.Funnels(),
		},
		client.Configure,
	)
}
