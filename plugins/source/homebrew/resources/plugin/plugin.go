package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/resources/services/analytics"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	tables := []*schema.Table{
		analytics.Installs30Days(),
		analytics.Installs90Days(),
		analytics.Installs365Days(),
	}
	return source.NewPlugin(
		"homebrew",
		Version,
		tables,
		client.Configure,
	)
}
