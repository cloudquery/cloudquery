package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/awspricing/client"
	"github.com/cloudquery/cloudquery/plugins/source/awspricing/resources"
	"github.com/cloudquery/plugin-sdk/v3/plugins/source"
	"github.com/cloudquery/plugin-sdk/v3/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"awspricing",
		Version,
		schema.Tables{
			resources.Services(),
		},
		client.New,
	)
}
