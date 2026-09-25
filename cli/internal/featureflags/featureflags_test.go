package featureflags

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

const testFlag = "cq-cli-test-flag"

func resetCache(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		fetchOnce = sync.Once{}
		cached = nil
	})
	fetchOnce = sync.Once{}
	cached = nil
}

func testContext() Context {
	return Context{UserID: "user-uuid", Team: "acme", Environment: "cli", CLIVersion: "v6.1.0"}
}

func serve(t *testing.T, handler http.HandlerFunc) {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	t.Setenv("CQ_LD_BASE_URL", server.URL)
	t.Setenv("CQ_LD_CLIENT_ID", "test-client-id")
}

func TestBoolFlagOn(t *testing.T) {
	resetCache(t)
	serve(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"` + testFlag + `":{"value":true,"version":3}}`))
	})

	if !BoolFlag(context.Background(), testFlag, false, testContext()) {
		t.Error("got false, want true")
	}
}

func TestBoolFlagDefaults(t *testing.T) {
	for _, tc := range []struct {
		name    string
		handler http.HandlerFunc
	}{
		{"flag absent", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"other-flag":{"value":true}}`))
		}},
		{"non-boolean value", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`{"` + testFlag + `":{"value":"yes"}}`))
		}},
		{"server error", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}},
		{"unparsable body", func(w http.ResponseWriter, _ *http.Request) {
			_, _ = w.Write([]byte(`not json`))
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resetCache(t)
			serve(t, tc.handler)

			if BoolFlag(context.Background(), testFlag, false, testContext()) {
				t.Error("got true, want the false default")
			}
			if !BoolFlag(context.Background(), testFlag, true, testContext()) {
				t.Error("got false, want the true default")
			}
		})
	}
}

func TestBoolFlagUnreachableEndpoint(t *testing.T) {
	resetCache(t)
	server := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	url := server.URL
	server.Close()
	t.Setenv("CQ_LD_BASE_URL", url)

	if BoolFlag(context.Background(), testFlag, false, testContext()) {
		t.Error("got true, want false when LaunchDarkly is unreachable")
	}
}

func TestBoolFlagWithoutUserSkipsRequest(t *testing.T) {
	resetCache(t)
	requests := 0
	serve(t, func(w http.ResponseWriter, _ *http.Request) {
		requests++
		_, _ = w.Write([]byte(`{"` + testFlag + `":{"value":true}}`))
	})

	if BoolFlag(context.Background(), testFlag, false, Context{}) {
		t.Error("got true, want false without a user")
	}
	if requests != 0 {
		t.Errorf("got %d requests, want 0", requests)
	}
}

func TestBoolFlagFetchesOncePerProcess(t *testing.T) {
	resetCache(t)
	requests := 0
	serve(t, func(w http.ResponseWriter, _ *http.Request) {
		requests++
		_, _ = w.Write([]byte(`{"` + testFlag + `":{"value":true}}`))
	})

	for range 3 {
		BoolFlag(context.Background(), testFlag, false, testContext())
	}
	if requests != 1 {
		t.Errorf("got %d requests, want 1", requests)
	}
}

func TestEvaluationContextIsSentInThePath(t *testing.T) {
	resetCache(t)
	var path string
	serve(t, func(w http.ResponseWriter, r *http.Request) {
		path = r.URL.Path
		_, _ = w.Write([]byte(`{}`))
	})

	BoolFlag(context.Background(), testFlag, false, testContext())

	prefix := "/sdk/evalx/test-client-id/contexts/"
	if !strings.HasPrefix(path, prefix) {
		t.Fatalf("got path %q, want prefix %q", path, prefix)
	}
	decoded, err := base64.URLEncoding.DecodeString(strings.TrimPrefix(path, prefix))
	if err != nil {
		t.Fatalf("failed to decode the context: %v", err)
	}
	var evalCtx map[string]string
	if err := json.Unmarshal(decoded, &evalCtx); err != nil {
		t.Fatalf("failed to unmarshal the context: %v", err)
	}
	for key, want := range map[string]string{
		"kind":        "user",
		"key":         "user-uuid",
		"team":        "acme",
		"environment": "cli",
		"cliVersion":  "v6.1.0",
	} {
		if evalCtx[key] != want {
			t.Errorf("context %q = %q, want %q", key, evalCtx[key], want)
		}
	}
}
