package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/resources/services/items"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	tables := []*schema.Table{
		items.Items(),
	}
	return source.NewPlugin(
		"hackernews",
		Version,
		tables,
		client.Configure,
	)
}
