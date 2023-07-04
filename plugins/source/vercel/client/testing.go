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

	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/scheduler"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/balancer/grpclb/state"
)

const testToken = "SomeToken"

type TestOptions struct {
	Backend state.State
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
	if req.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", testToken) {
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
		TeamIDs: []string{"test-team-id"},
	}
	spec.SetDefaults()

	if err := createServices(router); err != nil {
		t.Fatalf("failed to create services: %v", err)
	}

	services := vercel.New(l, mockClient, h.URL, testToken, spec.TeamIDs[0], 5, 10, 100)

	c := New(l, *spec, services, spec.TeamIDs, nil)

	messages, err := sched.SyncAll(context.Background(), c, schema.Tables{table})
	if err != nil {
		t.Fatalf("failed to sync: %v", err)
	}
	messages.InsertMessage()
	records := filterInserts(messages).GetRecordsForTable(table)
	emptyColumns := schema.FindEmptyColumns(table, records)
	if len(emptyColumns) > 0 {
		t.Fatalf("empty columns: %v", emptyColumns)
	}
}

func filterInserts(msgs message.SyncMessages) message.SyncInserts {
	inserts := []*message.SyncInsert{}
	for _, msg := range msgs {
		if m, ok := msg.(*message.SyncInsert); ok {
			inserts = append(inserts, m)
		}
	}
	return inserts
}
