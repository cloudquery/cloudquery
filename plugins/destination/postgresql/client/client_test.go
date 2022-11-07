package client

import (
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

func getTestLogger(t *testing.T) zerolog.Logger {
	t.Helper()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	return zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
}

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}

func TestPgPlugin(t *testing.T) {
	p := plugins.NewDestinationPlugin("postgresql", "development", New)
	plugins.DestinationPluginTestSuiteRunner(t, p, specs.Destination{
		WriteMode: specs.WriteModeOverwriteDeleteStale,
		Spec: Spec{
			ConnectionString: getTestConnection(),
			PgxLogLevel:      LogLevelTrace,
		},
	})
}
