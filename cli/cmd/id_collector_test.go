package cmd

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewIDCollectorSkipsUnmappedSource(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		_, _ = w.Write([]byte(`{"cq-cli-account-id-analytics":{"value":true}}`))
	}))
	defer server.Close()
	t.Setenv("CQ_LD_BASE_URL", server.URL)

	if collector := newIDCollector(context.Background(), "cloudquery/test"); collector != nil {
		t.Error("got a collector for an unmapped source, want nil")
	}
	if requests != 0 {
		t.Errorf("got %d flag requests, want 0 for an unmapped source", requests)
	}
}

func TestNewIDCollectorRequiresTelemetry(t *testing.T) {
	if collector := newIDCollector(context.Background(), "cloudquery/aws"); collector != nil {
		t.Error("got a collector with telemetry off, want nil")
	}
}
