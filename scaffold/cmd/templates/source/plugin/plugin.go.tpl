package plugin

import (
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/{{.Org}}/cq-source-{{.Name}}/client"
  "github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/{{.Org}}/cq-source-{{.Name}}/resources"
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
