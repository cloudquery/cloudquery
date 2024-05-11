package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/client"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/client/spec"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/resources/plugin"
	pluginSDK "github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://19d1257d36854a51b17c06614e76dc2d@o1396617.ingest.sentry.io/4503896817336320"
)

func main() {
	p := pluginSDK.NewPlugin(plugin.Name, plugin.Version, client.New,
		pluginSDK.WithKind(plugin.Kind),
		pluginSDK.WithTeam(plugin.Team),
		pluginSDK.WithJSONSchema(spec.JSONSchema),
	)
	server := serve.Plugin(p,
		serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server(),
	)

	done := instrumentPprof()
	defer done()
	err := server.Serve(context.Background())
	if err != nil {
		log.Println("failed to serve plugin:", err)
	}
}
