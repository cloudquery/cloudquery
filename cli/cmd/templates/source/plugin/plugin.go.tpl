package plugin

import (
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/{{.Org}}/cq-source-{{.Name}}/client"
	"github.com/{{.Org}}/cq-source-{{.Name}}/resources"
)

var (
	Version = "development"
)

func Plugin() *plugins.SourcePlugin {
	return plugins.NewSourcePlugin(
		"{{.Name}}",
		Version,
		schema.Tables{
			resources.SampleTable(),
		},
		client.New,
	)
}
