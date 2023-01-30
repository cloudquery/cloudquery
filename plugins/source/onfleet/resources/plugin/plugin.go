package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/onfleet/client"
	"github.com/cloudquery/cloudquery/plugins/source/onfleet/resources/services"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
)

var (
	Version = "development"
)

func Plugin() *source.Plugin {
	return source.NewPlugin(
		"cloudquery-onfleet",
		Version,
		schema.Tables{
			services.Hubs(),
			services.Workers(),
			services.Tasks(),
			services.Teams(),
		},
		client.New,
	)
}
