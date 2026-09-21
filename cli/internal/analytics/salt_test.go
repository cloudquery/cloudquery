package analytics

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func resetSaltCache(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		saltOnce = sync.Once{}
		cachedSalt = ""
	})
	saltOnce = sync.Once{}
	cachedSalt = ""
}

func TestTeamAnalyticsSaltWithoutTeam(t *testing.T) {
	resetSaltCache(t)
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requests++
		_, _ = w.Write([]byte(`{"salt":"6f1f0d2e-6b0a-4a1e-9a66-6a4a2f0d1c33"}`))
	}))
	defer server.Close()
	t.Setenv(envAPIURL, server.URL)

	if salt := TeamAnalyticsSalt(context.Background(), ""); salt != "" {
		t.Errorf("got salt %q, want empty without a team", salt)
	}
	if requests != 0 {
		t.Errorf("got %d requests, want 0 without a team", requests)
	}
}

func TestTeamAnalyticsSaltRequest(t *testing.T) {
	resetSaltCache(t)
	var path, authorization string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, authorization = r.URL.Path, r.Header.Get("Authorization")
		_, _ = w.Write([]byte(`{"salt":"6f1f0d2e-6b0a-4a1e-9a66-6a4a2f0d1c33"}`))
	}))
	defer server.Close()
	t.Setenv(envAPIURL, server.URL)
	t.Setenv("CLOUDQUERY_API_KEY", "cq_test_key")

	salt := TeamAnalyticsSalt(context.Background(), "acme")
	if TeamAnalyticsSalt(context.Background(), "acme") != salt {
		t.Error("the cached salt changed between calls")
	}
	if salt != "6f1f0d2e-6b0a-4a1e-9a66-6a4a2f0d1c33" {
		t.Errorf("got salt %q, want the served one", salt)
	}
	if path != "/teams/acme/analytics-salt" {
		t.Errorf("got path %q, want /teams/acme/analytics-salt", path)
	}
	if authorization != "Bearer cq_test_key" {
		t.Errorf("got authorization %q, want the bearer token", authorization)
	}
}

func TestTeamAnalyticsSaltFailures(t *testing.T) {
	for _, tc := range []struct {
		name    string
		handler http.HandlerFunc
	}{
		{"server error", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusInternalServerError) }},
		{"forbidden", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusForbidden) }},
		{"unparsable body", func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte(`not json`)) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resetSaltCache(t)
			server := httptest.NewServer(tc.handler)
			defer server.Close()
			t.Setenv(envAPIURL, server.URL)
			t.Setenv("CLOUDQUERY_API_KEY", "cq_test_key")

			if salt := TeamAnalyticsSalt(context.Background(), "acme"); salt != "" {
				t.Errorf("got salt %q, want empty", salt)
			}
		})
	}
}
