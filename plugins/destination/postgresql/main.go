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

func main() {
	p := pluginSDK.NewPlugin(plugin.Name, plugin.Version, client.New,
		pluginSDK.WithKind(plugin.Kind),
		pluginSDK.WithTeam(plugin.Team),
		pluginSDK.WithJSONSchema(spec.JSONSchema),
		pluginSDK.WithConnectionTester(client.ConnectionTester),
	)
	server := serve.Plugin(p,	serve.WithPluginSentryDSN(sentryDSN))

	done := instrumentPprof()
	defer done()
	err := server.Serve(context.Background())
	if err != nil {
		log.Println("failed to serve plugin:", err)
	}
}
