package client

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
)

func MockTestHelper(t *testing.T, table *schema.Table, createService func(*httprouter.Router) error) {
	version := "vDev"
	t.Helper()

	table.IgnoreInTests = false
	mux := httprouter.New()
	ts := httptest.NewUnstartedServer(mux)
	defer ts.Close()

	newTestExecutionClient := func(_ context.Context, logger zerolog.Logger, _ specs.Source, _ source.Options) (schema.ClientMeta, error) {
		err := createService(mux)
		if err != nil {
			return nil, fmt.Errorf("failed to createService: %w", err)
		}
		ts.Start()

		tsClient, err := (&Spec{
			APIKey:      "test-key",
			Tailnet:     "test-tailnet",
			EndpointURL: ts.URL,
		}).getClient()
		if err != nil {
			return nil, fmt.Errorf("failed to create client: %w", err)
		}

		c := &Client{
			Client:  tsClient,
			logger:  logger,
			tailnet: "test",
		}

		return c, nil
	}

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

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
