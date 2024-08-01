package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "k8s"
	Kind    = "source"
	Team    = "cloudquery"
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		Name,
		Version,
		newClient,
		plugin.WithJSONSchema(spec.JSONSchema),
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
	)
}
