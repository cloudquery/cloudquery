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

	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

const testToken = "SomeToken"

type MockHttpClient struct {
	rootURL string
	scheme  string
	host    string
	client  *http.Client
}

func NewMockHttpClient(cl *http.Client, rootURL string) *MockHttpClient {
	u, err := url.Parse(rootURL)
	if err != nil {
		panic(err)
	}
	return &MockHttpClient{
		client:  cl,
		rootURL: rootURL,
		scheme:  u.Scheme,
		host:    u.Host,
	}
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	if req.Header.Get("X-Shopify-Access-Token") != testToken {
		return &http.Response{StatusCode: http.StatusUnauthorized, Status: "401 Unauthorized", Body: io.NopCloser(bytes.NewReader(nil))}, nil
	}

	req.URL.Host = c.host
	req.URL.Scheme = c.scheme
	return c.client.Do(req)
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
	mockClient := NewMockHttpClient(h.Client(), h.URL)

	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	newTestExecutionClient := func(ctx context.Context, logger zerolog.Logger, spec specs.Source, _ source.Options) (schema.ClientMeta, error) {
		var spSpec Spec
		if err := spec.UnmarshalSpec(&spSpec); err != nil {
			return nil, fmt.Errorf("failed to unmarshal shopify spec: %w", err)
		}

		if err := createServices(router); err != nil {
			return nil, err
		}

		services, err := shopify.New(logger, mockClient, "", "", testToken, h.URL, 1)
		if err != nil {
			return nil, err
		}

		c := New(logger, spec, spSpec, services)
		return &c, nil
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
