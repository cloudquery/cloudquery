package client

import (
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins/destination"
)

func TestPlugin(t *testing.T) {
<<<<<<< HEAD
	p := destination.NewPlugin("snowflake", "development", New)
=======
	p := destination.NewPlugin("snowflake", "development", New, destination.WithManagerWriter())
>>>>>>> da4dc1c6c (feat(destinations): Migrate to managed batching SDK)
	destination.PluginTestSuiteRunner(t, p,
		Spec{
			// ConnectionString: os.Getenv("SNOW_TEST_DSN"),
			// REMBAZX.ZO72963
			ConnectionString: "yevgenyp:tur*jup@prm8ETC!uwk@yk49522.europe-west4.gcp/testdb/public?warehouse=test",
		},
		destination.PluginTestSuiteTests{
			SkipOverwrite: true,
		})
}
