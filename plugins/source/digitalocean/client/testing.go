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

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestHelper(t *testing.T, table *schema.Table, createService func(t *testing.T, ctrl *gomock.Controller) Services, options TestOptions) {
	t.Helper()
	table.IgnoreInTests = false
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		var doSpec Spec
		if err := spec.UnmarshalSpec(&doSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal do spec: %w", err)
		}
		l := zerolog.New(zerolog.NewTestWriter(t)).Output(
			zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
		).Level(zerolog.DebugLevel).With().Timestamp().Logger()

		ctrl := gomock.NewController(t)
		services := createService(t, ctrl)

		c := Client{
			logger:       l,
			SpacesRegion: "nyc3",
			Services:     &services,
		}
		return &c, nil
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
