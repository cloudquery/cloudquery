package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/xanzy/go-gitlab"
)

type TestOptions struct{}

func GitlabMockTestHelper(t *testing.T, table *schema.Table, createService func(*httprouter.Router) error, options TestOptions) {
	version := "vDev"
	t.Helper()

	table.IgnoreInTests = false
	mux := httprouter.New()
	ts := httptest.NewUnstartedServer(mux)
	defer ts.Close()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		err := createService(mux)
		if err != nil {
			return nil, fmt.Errorf("failed to createService: %w", err)
		}
		ts.Start()
		client, err := gitlab.NewClient("",
			gitlab.WithBaseURL(ts.URL),
			// Disable backoff to speed up tests that expect errors.
			gitlab.WithCustomBackoff(func(_, _ time.Duration, _ int, _ *http.Response) time.Duration {
				return 0
			}),
		)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}
		c := &Client{
			logger:  l,
			Gitlab:  client,
			BaseURL: ts.URL,
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

// func setup(t *testing.T) (*http.ServeMux, *gitlab.Client) {
// 	// mux is the HTTP request multiplexer used with the test server.
// 	mux := http.NewServeMux()

// 	// server is a test HTTP server used to provide mock API responses.
// 	server := httptest.NewServer(mux)
// 	t.Cleanup(server.Close)

// 	// client is the Gitlab client being tested.

// 	return mux, client
// }
