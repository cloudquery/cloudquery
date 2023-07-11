package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/plugins/destination/snowflake/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/snowflake/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://6640bc34f8d34a3d816f53d396fe997e@o1396617.ingest.sentry.io/4504208023224320"
)

func main() {
	p := plugin.NewPlugin("snowflake", internalPlugin.Version, client.New)
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
