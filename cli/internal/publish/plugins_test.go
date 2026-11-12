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

func TestValidateTokenRealm(t *testing.T) {
	tests := []struct {
		name           string
		realm          string
		allowPlainHTTP bool
		wantErr        bool
	}{
		{name: "https", realm: "https://auth.example.com/token"},
		{name: "http rejected", realm: "http://auth.example.com/token", wantErr: true},
		{name: "http allowed for local registry", realm: "http://localhost:5000/token", allowPlainHTTP: true},
		{name: "other scheme rejected for local registry", realm: "ftp://localhost/token", allowPlainHTTP: true, wantErr: true},
		{name: "missing scheme", realm: "auth.example.com/token", wantErr: true},
		{name: "unparsable", realm: "://bad", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTokenRealm(tt.realm, tt.allowPlainHTTP)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validateTokenRealm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsLocalhostRegistry(t *testing.T) {
	tests := []struct {
		domain string
		want   bool
	}{
		{domain: "localhost", want: true},
		{domain: "localhost:5000", want: true},
		{domain: "localhost.example.com", want: false},
		{domain: "localhost.example.com:5000", want: false},
		{domain: "localhostx", want: false},
		{domain: "docker.cloudquery.io", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.domain, func(t *testing.T) {
			if got := isLocalhostRegistry(tt.domain); got != tt.want {
				t.Fatalf("isLocalhostRegistry(%q) = %v, want %v", tt.domain, got, tt.want)
			}
		})
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
