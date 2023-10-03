package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/plugins/destination/sqlite/client"
	internalPlugin "github.com/cloudquery/cloudquery/plugins/destination/sqlite/resources/plugin"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/serve"
)

const (
	sentryDSN = "https://8e769078076443cd9c578833740beb54@o1396617.ingest.sentry.io/4504797525901312"
)

func main() {
	p := plugin.NewPlugin(
		"sqlite",
		internalPlugin.Version,
		client.New,
		plugin.WithBuildTargets([]plugin.BuildTarget{
			{OS: plugin.GoOSLinux, Arch: plugin.GoArchAmd64},
			{OS: plugin.GoOSWindows, Arch: plugin.GoArchAmd64},
			{OS: plugin.GoOSDarwin, Arch: plugin.GoArchAmd64},
			{OS: plugin.GoOSDarwin, Arch: plugin.GoArchArm64},
		}))
	if err := serve.Plugin(p, serve.WithPluginSentryDSN(sentryDSN), serve.WithDestinationV0V1Server()).Serve(context.Background()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
