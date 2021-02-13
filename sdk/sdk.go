package sdk

import (
	"github.com/cloudquery/cloudquery/cmd"
	"github.com/cloudquery/cloudquery/plugin"
	goplugin "github.com/hashicorp/go-plugin"
	"os"
)

func ServePlugin(provider plugin.CQProvider) {
	if len(os.Args) == 1 {
		goplugin.Serve(&goplugin.ServeConfig{
			HandshakeConfig: plugin.Handshake,
			VersionedPlugins: map[int]goplugin.PluginSet{
				1: {
					"provider": &plugin.CQPlugin{Impl: provider},
				}},

			// A non-nil value here enables gRPC serving for this plugin...
			GRPCServer: goplugin.DefaultGRPCServer,
		})
	} else {
		plugin.RegisterRunSelfProvider(provider)
		os.Args = append(os.Args, "--runself")
		cmd.Execute()
	}
}
