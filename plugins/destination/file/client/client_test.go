package client

import (
	"context"
	"testing"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

func newTestMode(ctx context.Context, logger zerolog.Logger, spec specs.Destination) (plugins.DestinationClient, error) {
	c, err := New(ctx, logger, spec)
	if err != nil {
		return nil, err
	}
	c.(*Client).testMode = true
	return c, nil
}

func TestPluginCSV(t *testing.T) {
	p := plugins.NewDestinationPlugin("file", "development", newTestMode)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: t.TempDir(),
			Backend:   BackendTypeLocal,
			Format:    FormatTypeCSV,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
		},
	)

	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: "cq-yev-test/dest-plugin-file",
			Backend:   BackendTypeGCS,
			Format:    FormatTypeCSV,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
			SkipAppendTwice: true,
		},
	)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: "cq-playground-test/dest-plugin-file",
			Backend:   BackendTypeS3,
			Format:    FormatTypeCSV,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
			SkipAppendTwice: true,
		},
	)
}

func TestPluginJSON(t *testing.T) {
	p := plugins.NewDestinationPlugin("file", "development", newTestMode)
	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: t.TempDir(),
			Backend:   BackendTypeLocal,
			Format:    FormatTypeJSON,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
		},
	)

	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: "cq-yev-test/dest-plugin-file",
			Backend:   BackendTypeGCS,
			Format:    FormatTypeJSON,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
			SkipAppendTwice: true,
		},
	)

	plugins.DestinationPluginTestSuiteRunner(t, p,
		Spec{
			Directory: "cq-playground-test/dest-plugin-file",
			Backend:   BackendTypeS3,
			Format:    FormatTypeJSON,
		},
		plugins.DestinationTestSuiteTests{
			SkipOverwrite:   true,
			SkipDeleteStale: true,
			SkipAppendTwice: true,
		},
	)
}
