package cmd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cloudquery/cli/v6/internal/analytics"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
)

var fileDestination = []specs.Destination{{Metadata: specs.Metadata{Name: "file"}}}

func TestNewIDCollectorSkipsUnmappedSource(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		_, _ = w.Write([]byte(`{"cq-cli-account-id-analytics":{"value":true}}`))
	}))
	defer server.Close()
	t.Setenv("CQ_LD_BASE_URL", server.URL)

	if _, ok := newIDCollector(context.Background(), "cloudquery/test", fileDestination).(analytics.NoopIDCollector); !ok {
		t.Error("got a collecting collector for an unmapped source, want the no-op one")
	}
	if requests != 0 {
		t.Errorf("got %d flag requests, want 0 for an unmapped source", requests)
	}
}

func TestNewIDCollectorRequiresTelemetry(t *testing.T) {
	if _, ok := newIDCollector(context.Background(), "cloudquery/aws", fileDestination).(analytics.NoopIDCollector); !ok {
		t.Error("got a collecting collector with telemetry off, want the no-op one")
	}
}

func TestNewIDCollectorSkipsPlatformOnlySyncs(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		_, _ = w.Write([]byte(`{"cq-cli-account-id-analytics":{"value":true}}`))
	}))
	defer server.Close()
	t.Setenv("CQ_LD_BASE_URL", server.URL)

	platformOnly := []specs.Destination{{Metadata: specs.Metadata{Name: "platform"}}}
	if _, ok := newIDCollector(context.Background(), "cloudquery/aws", platformOnly).(analytics.NoopIDCollector); !ok {
		t.Error("got a collecting collector for a platform-only sync, want the no-op one")
	}
	if requests != 0 {
		t.Errorf("got %d flag requests, want 0 for a platform-only sync", requests)
	}
}
