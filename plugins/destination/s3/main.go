package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/plugins/destination/s3/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/s3/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://c808d26943414033b2fb8bb5b5822ab9@o1396617.ingest.sentry.io/4504407917592576"
)

func main() {
	p := plugin.NewPlugin("s3", internalPlugin.Version, client.New)
	serveHelper(serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve)
}

func serveHelper(f func(context.Context) error) {
	if err := f(context.Background()); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
