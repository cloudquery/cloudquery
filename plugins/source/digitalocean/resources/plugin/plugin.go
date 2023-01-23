package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"digitalocean",
		Version,
		Tables(),
		client.New,
	)
}
