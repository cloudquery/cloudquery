package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/file/v5/client"
	"github.com/cloudquery/cloudquery/plugins/destination/file/v5/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/file/v5/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://1e33dfd084aa43f2aa8e686f15a64e45@o1396617.ingest.us.sentry.io/4504407264526336"
)

func main() {
	p := plugin.NewPlugin(
		internalPlugin.Name,
		internalPlugin.Version,
		client.New,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(spec.JSONSchema),
	)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
