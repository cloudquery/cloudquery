package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec any) (plugin.Client, error) {
		return &Client{
			logger:   l,
			Homebrew: builder(t),
		}, nil
	}
	p := plugin.NewPlugin(
		table.Name,
		version,
		newTestExecutionClient,
	)
	p.SetLogger(l)
	plugin.TestPluginSync(t, p, &Spec{}, plugin.SyncOptions{
		Tables:            []string{table.Name},
		SkipTables:        nil,
		Concurrency:       0,
		DeterministicCQID: false,
		StateBackend:      nil,
	})
}
