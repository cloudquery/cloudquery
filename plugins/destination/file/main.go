package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/plugins/destination/file/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/file/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://1e33dfd084aa43f2aa8e686f15a64e45@o1396617.ingest.sentry.io/4504407264526336"
)

func main() {
	p := plugin.NewPlugin("file", internalPlugin.Version, client.New)
	serveHelper(serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve)
}

func serveHelper(f func(context.Context) error) {
	if err := f(context.Background()); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
