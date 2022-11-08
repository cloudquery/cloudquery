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
	version := "vDev"

	t.Helper()

	table.IgnoreInTests = false
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	newTestExecutionClient := func(ctx context.Context, _ zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		svc, err := createService()
		if err != nil {
			return nil, fmt.Errorf("failed to createService %w", err)
		}
		var herokuSpec Spec
		if err := spec.UnmarshalSpec(&herokuSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
		}
		c := &Client{
			logger: l,
			Heroku: svc,
		}

		return c, nil
	}

	p := plugins.NewSourcePlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient)
	p.SetLogger(l)
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
