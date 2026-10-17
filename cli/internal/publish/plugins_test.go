package publish

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/distribution/reference"
)

func TestGetDockerTokenTokenRequestError(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", "Bearer realm=\"://bad\",service=\"registry\",scope=\"repository:test:pull\"")
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer ts.Close()

	host := strings.TrimPrefix(ts.URL, "https://")
	ref, err := reference.ParseNormalizedNamed(fmt.Sprintf("%s/test", host))
	if err != nil {
		t.Fatalf("failed to parse reference: %v", err)
	}

	_, err = getDockerToken(context.Background(), ref, "latest", "team", "user", "pass", true)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	if !strings.Contains(err.Error(), "could not create request") {
		t.Fatalf("expected a %q error, got: %v", "could not create request", err)
	}
}

func TestGetSpecJsonScheme(t *testing.T) {
	dir := t.TempDir()

	got, err := GetSpecJsonScheme(dir)
	if err != nil {
		t.Fatalf("unexpected error for missing file: %v", err)
	}
	if got != nil {
		t.Fatalf("expected nil for missing file, got: %v", *got)
	}

	want := `{"type":"object"}`
	if err := os.WriteFile(filepath.Join(dir, "spec_json_schema.json"), []byte(want), 0o644); err != nil {
		t.Fatalf("failed to write spec_json_schema.json: %v", err)
	}

	got, err = GetSpecJsonScheme(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || *got != want {
		t.Fatalf("got %v, want %q", got, want)
	}
}
