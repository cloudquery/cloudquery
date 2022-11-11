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
	"github.com/rs/zerolog"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestHelper(t *testing.T, table *schema.Table, createService func() (*Services, error), options TestOptions) {
	version := "vDev"
	t.Helper()

	table.IgnoreInTests = false

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		svc, err := createService()
		if err != nil {
			return nil, fmt.Errorf("failed to createService: %w", err)
		}
		var gcpSpec Spec
		if err := spec.UnmarshalSpec(&gcpSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal gcp spec: %w", err)
		}
		c := &Client{
			// plugin: p,
			logger: logger,
			// logger:   t.Log(),
			Services: svc,
			projects: []string{"testProject"},
		}

		return c, nil
	}
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

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
