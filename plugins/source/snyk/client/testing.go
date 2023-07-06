package client

import (
	"context"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-pb-go/specs"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/julienschmidt/httprouter"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func MockTestHelper(t *testing.T, table *schema.Table, createService func(*httprouter.Router) error) {
	version := "vDev"
	t.Helper()

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))

	mux := httprouter.New()
	ts := httptest.NewUnstartedServer(mux)
	defer ts.Close()

	spec := &Spec{
		APIKey:        "test-key",
		Organizations: []snyk.Organization{
			{ID: "test-org-id", Name: "test-org-name", Group: &snyk.Group{
				ID:   "test-group-id",
				Name: "test-group-name",
			}},
		EndpointURL:   ts.URL + "/",
	}
	spec.SetDefaults()
	require.NoError(t, spec.Validate())
	if err := spec.Validate(); err != nil {
		t.Fatalf("failed to validate spec: %v", err)
	}
	table.IgnoreInTests = false

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, _ specs.Source, _ source.Options) (schema.ClientMeta, error) {
		err := createService(mux)
		if err != nil {
			return nil, fmt.Errorf("failed to createService: %w", err)
		}
		ts.Start()

		snykClient := snyk.NewClient("test-key", snyk.WithBaseURL(ts.URL+"/"))
		if err != nil {
			return nil, fmt.Errorf("failed to create client: %w", err)
		}

		c := &Client{
			Client: snykClient,
			logger: logger,
			Organizations:
			},
		}
		return c, nil
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
