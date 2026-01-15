package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/sqlite/v2/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/sqlite/v2/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://8e769078076443cd9c578833740beb54@o1396617.ingest.us.sentry.io/4504797525901312"
)

func main() {
	p := plugin.NewPlugin(
		internalPlugin.Name,
		internalPlugin.Version,
		client.New,
		plugin.WithBuildTargets([]plugin.BuildTarget{
			{OS: plugin.GoOSLinux, Arch: plugin.GoArchAmd64, CGO: true},
			{OS: plugin.GoOSLinux, Arch: plugin.GoArchArm64, CGO: true, Env: []string{"CC=aarch64-linux-gnu-gcc"}},
			{OS: plugin.GoOSWindows, Arch: plugin.GoArchAmd64, CGO: true},
			{OS: plugin.GoOSDarwin, Arch: plugin.GoArchAmd64, CGO: true},
			{OS: plugin.GoOSDarwin, Arch: plugin.GoArchArm64, CGO: true},
		}),
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(client.JSONSchema),
	)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
