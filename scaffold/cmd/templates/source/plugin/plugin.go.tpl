package plugin

import (
	"github.com/{{.Org}}/cq-source-{{.Name}}/client"
	"github.com/{{.Org}}/cq-source-{{.Name}}/resources"

	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"{{.Org}}-{{.Name}}",
		Version,
		schema.Tables{
			resources.SampleTable(),
		},
		client.New,
	)
}
