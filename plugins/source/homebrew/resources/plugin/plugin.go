package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/resources/services/analytics"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	tables := []*schema.Table{
		analytics.Installs(homebrew.Days30),
		analytics.Installs(homebrew.Days90),
		analytics.Installs(homebrew.Days365),
		analytics.CaskInstalls(homebrew.Days30),
		analytics.CaskInstalls(homebrew.Days90),
		analytics.CaskInstalls(homebrew.Days365),
		analytics.BuildErrors(homebrew.Days30),
		analytics.BuildErrors(homebrew.Days90),
		analytics.BuildErrors(homebrew.Days365),
	}
	return source.NewPlugin(
		"homebrew",
		Version,
		tables,
		client.Configure,
	)
}
