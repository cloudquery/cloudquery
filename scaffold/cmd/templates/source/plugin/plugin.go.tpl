package plugin

import (
	"github.com/{{.Org}}/cq-source-{{.Name}}/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin("{{.Org}}-{{.Name}}", Version, client.New)
}
