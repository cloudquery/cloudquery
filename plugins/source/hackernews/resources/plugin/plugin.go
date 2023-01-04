package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/resources/services/items"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
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
