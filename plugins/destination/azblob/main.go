package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/azblob/v4/client"
	"github.com/cloudquery/cloudquery/plugins/destination/azblob/v4/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/azblob/v4/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://43b48b4844804de7aebffe352b044f2c@o1396617.ingest.sentry.io/4504411507392512"
)

func main() {
	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(spec.JSONSchema),
	)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN),
		serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatalf("failed to serve plugin: %v", err)
	}
}
