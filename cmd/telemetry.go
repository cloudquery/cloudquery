package cmd

import (
	"github.com/cloudquery/cloudquery/internal/telemetry"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

var tele *telemetry.Client

func initTelemetry() {
	var opts []telemetry.Option

	if viper.GetBool("no-telemetry") {
		opts = append(opts, telemetry.WithDisabled())
	}

	if viper.GetBool("inspect-telemetry") {
		fs := afero.NewOsFs()
		f, err := fs.Create("cq-telemetry.txt")
		if err != nil {
			panic(err)
		}
		opts = append(opts, telemetry.WithExporter(f))
	}

	tele = telemetry.New(opts...)
}
