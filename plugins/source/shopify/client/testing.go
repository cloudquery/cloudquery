package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/state"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

const testToken = "SomeToken"

type TestOptions struct {
	Backend state.Client
}

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

func MockTestHelper(t *testing.T, table *schema.Table, createServices func(*mux.Router) error, _ TestOptions) {
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

	l := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()
	sched := scheduler.NewScheduler(scheduler.WithLogger(l))

	spec := &Spec{
		AccessToken: testToken,
		ShopURL:     "https://test.myshopify.com",
	}
	spec.SetDefaults()
	if err := spec.Validate(); err != nil {
		t.Fatalf("failed to validate spec: %v", err)
	}

	if err := createServices(router); err != nil {
		t.Fatalf("failed to create services: %v", err)
	}

	services, err := shopify.New(shopify.ClientOptions{
		Log:         l,
		HC:          mockClient,
		AccessToken: spec.AccessToken,
		ShopURL:     h.URL,
		MaxRetries:  1,
		PageSize:    50,
	})
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	c := New(l, *spec, services, nil)

	messages, err := sched.SyncAll(context.Background(), c, schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	records := messages.GetInserts().GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}
