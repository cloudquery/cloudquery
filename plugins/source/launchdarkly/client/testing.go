package client

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/gorilla/mux"
	ldapi "github.com/launchdarkly/api-client-go/v11"
	"github.com/rs/zerolog"
)

const (
	testToken = "some_token"
)

type TestOptions struct {
	Backend backend.Backend
}

func MockTestHelper(t *testing.T, table *schema.Table, createServices func(*mux.Router) error, opts TestOptions) {
	version := "vDev"

	t.Helper()
	table.IgnoreInTests = false

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Router received request to %s", r.URL.String())
		http.Error(w, "not found", http.StatusNotFound)
	})
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Authorization") != testToken {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	h := httptest.NewServer(router)
	defer h.Close()

	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		var ldSpec Spec
		if err := spec.UnmarshalSpec(&ldSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal launchdarkly spec: %w", err)
		}

		if err := createServices(router); err != nil {
			return nil, err
		}

		u, err := url.Parse(h.URL)
		if err != nil {
			return nil, err
		}

		cfg := ldapi.NewConfiguration()
		cfg.Host = u.Host
		cfg.Scheme = u.Scheme
		cfg.AddDefaultHeader("Authorization", testToken)
		cfg.HTTPClient = &http.Client{
			Timeout: 1 * time.Second,
		}
		services := ldapi.NewAPIClient(cfg)

		c := New(logger, spec, ldSpec, services, opts.Backend)
		return c, nil
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
