package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

func MockTestHelper(t *testing.T, table *schema.Table, createServices func(t *testing.T, ctrl *gomock.Controller) services.Services) {
	t.Helper()

	table.IgnoreInTests = false

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source) (schema.ClientMeta, error) {
		svc := createServices(t, gomock.NewController(t))
		l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro}).Level(zerolog.DebugLevel).With().Timestamp().Logger()
		servicesMap := make(map[string]*services.Services)
		servicesMap["testSubscription"] = &svc
		c := &Client{
			logger:        l,
			services:      servicesMap,
			subscriptions: []string{"testSubscription"},
		}

		return c, nil
	}

	p := plugins.NewSourcePlugin(table.Name, "dev", []*schema.Table{table}, newTestExecutionClient)
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:   "dev",
		Tables: []string{table.Name},
	})
}
