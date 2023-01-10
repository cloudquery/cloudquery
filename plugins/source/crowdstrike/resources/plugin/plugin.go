package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/crowdstrike/client"
	"github.com/cloudquery/cloudquery/plugins/source/crowdstrike/resources/services/alerts"
	"github.com/cloudquery/cloudquery/plugins/source/crowdstrike/resources/services/incidents"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	allTables := []*schema.Table{
		incidents.Crowdscore(),
		alerts.Query(),
	}
	return source.NewPlugin(
		"crowdstrike",
		Version,
		allTables,
		client.New,
	)
}
