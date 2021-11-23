package cmd

import (
	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

func telemetryOpts() []telemetry.Option {
	opts := []telemetry.Option{
		telemetry.WithVersionInfo(client.Version, Commit, Date),
		telemetry.WithLogger(hclog.Default()),
		telemetry.WithEndpoint(viper.GetString("telemetry-endpoint"), viper.GetBool("insecure-telemetry-endpoint")),
	}

	if viper.GetBool("no-telemetry") {
		opts = append(opts, telemetry.WithDisabled())
	}

	if viper.GetBool("debug-telemetry") {
		opts = append(opts, telemetry.WithDebug())
	}

	if viper.GetBool("inspect-telemetry") {
		fs := afero.NewOsFs()
		f, err := fs.Create("cq-telemetry.txt")
		if err != nil {
			panic(err)
		}
		opts = append(opts, telemetry.WithExporter(f))
	}

	return opts
}
