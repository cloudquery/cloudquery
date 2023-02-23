package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

type TestOptions struct{}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T) *homebrew.Client, opts TestOptions) {
	version := "vDev"
	table.IgnoreInTests = false
	t.Helper()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		return &Client{
			logger:   l,
			Homebrew: builder(t),
			Spec:     Spec{},
		}, nil
	}
	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	p.SetLogger(l)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
