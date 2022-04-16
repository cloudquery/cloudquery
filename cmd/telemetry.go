package cmd

import (
	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/ui"

	"github.com/hashicorp/go-hclog"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

func telemetryOpts() []telemetry.Option {
	opts := []telemetry.Option{
		telemetry.WithVersionInfo(core.Version, Commit, Date),
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
		const fn = "cq-telemetry.txt"
		fs := afero.NewOsFs()
		f, err := fs.Create(fn)
		if err != nil {
			panic(err)
		}
		ui.ColorizedOutput(ui.ColorInfo, "Created %s file locally\n", fn)
		opts = append(opts, telemetry.WithExporter(f))
	}

	return opts
}
