//go:build linux && boringcrypto

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
		plugin.WithJSONSchema(client.JSONSchema),
		plugin.WithConnectionTester(TestConnection),
		plugin.WithBuildTargets([]plugin.BuildTarget{
			{OS: plugin.GoOSLinux, Arch: plugin.GoArchAmd64, CGO: true, IncludeSymbols: true},
			{OS: plugin.GoOSLinux, Arch: plugin.GoArchArm64, CGO: true, IncludeSymbols: true},
		}),
	)
}
