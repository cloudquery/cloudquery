package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/motherduck/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/motherduck/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://d4bc6f4eb3014b8994c7a66846d86e18@o1396617.ingest.sentry.io/4504797281779712"
)

func main() {
	p := plugin.NewPlugin(
		internalPlugin.Name,
		internalPlugin.Version,
		client.New,
		plugin.WithBuildTargets([]plugin.BuildTarget{
			{OS: plugin.GoOSLinux, Arch: plugin.GoArchAmd64, CGO: true},
			{OS: plugin.GoOSDarwin, Arch: plugin.GoArchAmd64, CGO: true},
			{OS: plugin.GoOSDarwin, Arch: plugin.GoArchArm64, CGO: true},
		}),
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(client.JSONSchema),
	)
	server := serve.Plugin(p,
		serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server(),
	)
	err := server.Serve(context.Background())
	if err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
