package client

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

const (
	TestAccountID = "test_account"
	TestZoneID    = "test_zone"
)

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Clients) {
	version := "vDev"

	t.Helper()
	table.IgnoreInTests = false

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	newTestExecutionClient := func(ctx context.Context, _ zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		var cfSpec Spec
		if err := spec.UnmarshalSpec(&cfSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cloudflare spec: %w", err)
		}

		clients := builder(t, ctrl)
		c := New(logger, clients, clients[TestAccountID], AccountZones{
			TestAccountID: {
				AccountId: TestAccountID,
				Zones:     []string{TestZoneID},
			},
		})

		return c.withZoneID(TestAccountID, TestZoneID), nil
	}

	p := source.NewPlugin(
		table.Name,
		version,
		[]*schema.Table{
			table,
		},
		newTestExecutionClient,
	)
	p.SetLogger(logger)
	source.TestPluginSync(t, p, specs.Source{
		Name:         "dev",
		Path:         "cloudquery/dev",
		Version:      version,
		Tables:       []string{table.Name},
		Destinations: []string{"mock-destination"},
	})
}
