package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/awspricing/client"
	"github.com/cloudquery/cloudquery/plugins/source/awspricing/resources"
	"github.com/cloudquery/plugin-sdk/v2/plugins/source"
	"github.com/cloudquery/plugin-sdk/v2/schema"
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
