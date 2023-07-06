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

	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
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
	sched := scheduler.NewScheduler(scheduler.WithLogger(logger))

	spec := &Spec{
		Token:  testToken,
		Domain: "https://example.com",
	}
	spec.SetDefaults(&logger)
	if err := spec.Validate(); err != nil {
		t.Fatalf("failed to validate spec: %v", err)
	}

	if err := createServices(router); err != nil {
		t.Fatalf("failed to create services: %v", err)
	}

	cf := okta.NewConfiguration(
		okta.WithOrgUrl(h.URL),
		okta.WithToken(spec.Token),
		okta.WithCache(true),
		okta.WithTestingDisableHttpsCheck(true),
	)
	cf.HTTPClient = h.Client()
	cf.HTTPClient.Transport = &rt{
		RewriteBaseURL: h.URL,
		ParentRT:       http.DefaultTransport,
	}

	c := New(logger, *spec, okta.NewAPIClient(cf))

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
