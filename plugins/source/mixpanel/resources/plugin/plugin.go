package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/resources/services/annotations"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/resources/services/cohorts"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/resources/services/engage"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/resources/services/export"
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
			annotations.Annotations(),
			cohorts.Cohorts(),
			funnels.Funnels(),
			engage.EngageRevenues(),
			export.ExportEvents(),
		},
		client.Configure,
	)
}
