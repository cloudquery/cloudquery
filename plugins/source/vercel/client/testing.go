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
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) VercelServices) {
	version := "vDev"

	t.Helper()
	table.IgnoreInTests = false

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	newTestExecutionClient := func(ctx context.Context, _ zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		var veSpec Spec
		if err := spec.UnmarshalSpec(&veSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal vercel spec: %w", err)
		}

		services := builder(t, ctrl)
		c := New(logger, spec, veSpec, services, nil)
		return &c, nil
	}

	p := plugins.NewSourcePlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient,
	)
	p.SetLogger(logger)
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
