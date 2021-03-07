package sdk

import (
	"log"
	"os"

	"github.com/cloudquery/cloudquery/cmd"
	"github.com/cloudquery/cloudquery/plugin"
	"github.com/hashicorp/go-hclog"
	goplugin "github.com/hashicorp/go-plugin"
)

type ServeOpts struct {
	// Required: Name of provider.
	Name string

	// Required: Provider is the actual provider that will be served.
	Provider plugin.CQProvider

	// Optional: Logger is the logger that go-plugin will use.
	Logger hclog.Logger

	// Optional: Set NoLogOutputOverride to not override the log output with an hclog
	// adapter. This should only be used when running the plugin in
	// acceptance tests.
	NoLogOutputOverride bool
}

func ServePlugin(opts ServeOpts) {

	if opts.Name == "" {
		panic("missing provider name")
	}

	if opts.Provider == nil {
		panic("missing provider instance")
	}

	// Check of CQ_PROVIDER_DEBUG is turned on. In case it's true we self register the plugin and execute
	// CloudQuery main command line accepting any args same as the main binary. The client will execute the
	// only this plugin instead of using the downloaded provider plugins
	if os.Getenv("CQ_PROVIDER_DEBUG") != "" {
		plugin.GetManager().AddEmbeddedPlugin(opts.Name, opts.Provider)
		cmd.Execute()
		return
	}

	if !opts.NoLogOutputOverride {
		// In order to allow go-plugin to correctly pass log-levels through to
		// cloudquery, we need to use an hclog.Logger with JSON output. We can
		// inject this into the std `log` package here, so existing providers will
		// make use of it automatically.
		logger := hclog.New(&hclog.LoggerOptions{
			// We send all output to CloudQuery. Go-plugin will take the output and
			// pass it through another hclog.Logger on the client side where it can
			// be filtered.
			Level:      hclog.Trace,
			JSONFormat: true,
		})
		log.SetOutput(logger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true}))
	}
	goplugin.Serve(&goplugin.ServeConfig{
		HandshakeConfig: plugin.Handshake,
		VersionedPlugins: map[int]goplugin.PluginSet{
			1: {
				"provider": &plugin.CQPlugin{Impl: opts.Provider},
			}},
		GRPCServer: goplugin.DefaultGRPCServer,
		Logger:     opts.Logger,
	})
}
