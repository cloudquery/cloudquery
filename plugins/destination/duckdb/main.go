package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/duckdb/v5/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/duckdb/v5/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	p := plugin.NewPlugin(
		internalPlugin.Name,
		internalPlugin.Version,
		client.New,
		plugin.WithBuildTargets([]plugin.BuildTarget{
			{OS: plugin.GoOSLinux, Arch: plugin.GoArchAmd64, CGO: true},
			{OS: plugin.GoOSLinux, Arch: plugin.GoArchArm64, CGO: true},
			{OS: plugin.GoOSDarwin, Arch: plugin.GoArchAmd64, CGO: true},
			{OS: plugin.GoOSDarwin, Arch: plugin.GoArchArm64, CGO: true},
		}),
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(client.JSONSchema),
		plugin.WithConnectionTester(client.TestConnection),
	)

	if err := serve.Plugin(p, serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
