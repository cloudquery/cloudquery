package versions

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_GetLatestProviderRelease(t *testing.T) {
	cloudQueryServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/source-aws.json" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, `{"latest":"plugins/source/test/v1.2.3"}`)
	}))
	defer cloudQueryServer.Close()

	githubServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/some-org/cq-target-postgres/releases/latest" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, `{"tag_name":"v4.5.6"}`)
	}))
	defer githubServer.Close()

	c := NewClient()
	c.cloudQueryBaseURL = cloudQueryServer.URL
	c.githubBaseURL = githubServer.URL

	ctx := context.Background()
	version, err := c.GetLatestProviderRelease(ctx, CloudQueryOrg, "source", "aws")
	if err != nil {
		t.Fatalf("error calling GetLatestProviderRelease: %v", err)
	}
	if version != "v1.2.3" {
		t.Errorf("got cloudquery org version = %q, want %q", version, "v1.2.3")
	}

	githubVersion, err := c.GetLatestProviderRelease(ctx, "some-org", "target", "postgres")
	if err != nil {
		t.Fatalf("error calling GetLatestProviderRelease: %v", err)
	}
	if githubVersion != "v4.5.6" {
		t.Errorf("got community provider version = %q, want %q", version, "v4.5.6")
	}
}
