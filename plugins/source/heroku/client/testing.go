package client

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/rs/zerolog"
)

type TestOptions struct{}

func MockTestHelper(t *testing.T, table *schema.Table, createService func() (*heroku.Service, error), options TestOptions) {
	t.Helper()

	table.IgnoreInTests = false

	newTestExecutionClient := func(ctx context.Context, _ zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		svc, err := createService()
		if err != nil {
			return nil, fmt.Errorf("failed to createService %w", err)
		}
		var herokuSpec Spec
		if err := spec.UnmarshalSpec(&herokuSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
		}
		l := zerolog.New(zerolog.NewTestWriter(t)).Output(
			zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
		).Level(zerolog.DebugLevel).With().Timestamp().Logger()
		c := &Client{
			logger: l,
			Heroku: svc,
		}

		return c, nil
	}

	p := plugins.NewSourcePlugin(
		table.Name,
		"dev",
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:   "dev",
		Tables: []string{table.Name},
	})
}
