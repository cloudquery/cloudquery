package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	{{range .Packages}}"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/{{.}}"
    {{end}}"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "Development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"azure",
		Version,
		[]*schema.Table{
			{{range .Resources}}{{.}}(),
            {{end}}
		},
		client.New,
	)
}
