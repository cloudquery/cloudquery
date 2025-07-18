//go:build !fipsEnabled

package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/test/v4/client"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
)

var (
	Name    = "test"
	Kind    = "source"
	Team    = "cloudquery"
	Version = "development"
)

func Plugin() *plugin.Plugin {
	return plugin.NewPlugin(
		Name,
		Version,
		Configure,
		plugin.WithKind(Kind),
		plugin.WithTeam(Team),
		plugin.WithBuildTargets(buildTargets()),
		plugin.WithJSONSchema(client.JSONSchema),
		plugin.WithConnectionTester(TestConnection),
	)
}
