package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/gorilla/mux"
	"github.com/okta/okta-sdk-golang/v3/okta"
	"github.com/rs/zerolog"
)

const testToken = "SomeToken"

type rt struct {
	RewriteBaseURL string
	ParentRT       http.RoundTripper
}

func (r *rt) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get("Authorization") != fmt.Sprintf("SSWS %s", testToken) {
		return &http.Response{StatusCode: http.StatusUnauthorized, Status: "401 Unauthorized", Body: io.NopCloser(bytes.NewReader(nil))}, nil
	}

	if r.RewriteBaseURL != "" {
		u, err := url.Parse(r.RewriteBaseURL)
		if err != nil {
			return nil, err
		}
		req.URL.Host = u.Host
	}

	return r.ParentRT.RoundTrip(req)
}

func MockTestHelper(t *testing.T, table *schema.Table, createServices func(*mux.Router) error) {
	version := "vDev"

	t.Helper()
	table.IgnoreInTests = false

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Router received request to %s", r.URL.String())
		http.Error(w, "not found", http.StatusNotFound)
	})

	h := httptest.NewServer(router)
	defer h.Close()

	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	newTestExecutionClient := func(ctx context.Context, _ zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		var clientSpec Spec
		if err := spec.UnmarshalSpec(&clientSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal client spec: %w", err)
		}

		if err := createServices(router); err != nil {
			return nil, err
		}

		cf := okta.NewConfiguration(
			okta.WithOrgUrl(h.URL),
			okta.WithToken(testToken),
			okta.WithCache(true),
			okta.WithTestingDisableHttpsCheck(true),
		)
		cf.HTTPClient = h.Client()
		cf.HTTPClient.Transport = &rt{
			RewriteBaseURL: h.URL,
			ParentRT:       http.DefaultTransport,
		}

		return New(logger, spec, okta.NewAPIClient(cf)), nil
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
