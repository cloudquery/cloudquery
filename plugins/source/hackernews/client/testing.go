package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client/services"
	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

type TestOptions struct {
	Backend   backend.Backend
	StartTime time.Time
}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) services.HackernewsClient, opts TestOptions) {
	version := "vDev"
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		startTime := ""
		if !opts.StartTime.IsZero() {
			startTime = opts.StartTime.Format(time.RFC3339)
		}
		return &Client{
			logger:     l,
			Backend:    opts.Backend,
			HackerNews: builder(t, ctrl),
			Spec:       Spec{ItemConcurrency: 10, StartTime: startTime},
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
