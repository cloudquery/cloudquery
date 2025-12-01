//go:build fipsEnabled

//go:debug fips140=on

package main

import (
	"context"
	"crypto/fips140"
	"log"

	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client"
	"github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/client/spec"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/postgresql/v8/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

func main() {
	log.Printf("FIPS enabled: %t", fips140.Enabled())

	p := plugin.NewPlugin(internalPlugin.Name, internalPlugin.Version, client.New,
		plugin.WithBuildTargets(buildTargets()),
		plugin.WithKind(internalPlugin.Kind),
		plugin.WithTeam(internalPlugin.Team),
		plugin.WithJSONSchema(spec.JSONSchema),
		plugin.WithConnectionTester(client.ConnectionTester),
	)
	if err := serve.Plugin(p, serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func buildTargets() []plugin.BuildTarget {
	// default build targets for FIPS builds
	// fipsEnabled is used to enable FIPS mode
	targets := make([]plugin.BuildTarget, len(plugin.DefaultBuildTargets))
	for i := range plugin.DefaultBuildTargets {
		targets[i] = plugin.DefaultBuildTargets[i]
		targets[i].Tags = append(targets[i].Tags, "fipsEnabled")
	}
	return targets
}
