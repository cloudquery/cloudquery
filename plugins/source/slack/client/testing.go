package client

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/slack/client/services"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
)

type TestOptions struct{}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) services.SlackClient, _ TestOptions) {
	version := "vDev"
	table.IgnoreInTests = false
	t.Helper()
	ctrl := gomock.NewController(t)
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		return &Client{
			logger: l,
			Slack:  builder(t, ctrl),
			Teams:  []slack.Team{{ID: "test_team"}},
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
