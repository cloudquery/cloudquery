package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAddonPublish(t *testing.T) {
	t.Setenv("CLOUDQUERY_API_KEY", "testkey")

	wantCalls := map[string]int{
		"PUT /addons/cloudquery/test/versions/v1.2.3":         1,
		"POST /addons/cloudquery/test/versions/v1.2.3/assets": 1,
		"PUT /upload-zip": 1,
	}
	gotCalls := map[string]int{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		gotCalls[r.Method+" "+r.URL.Path]++
		switch r.URL.Path {
		case "/addons/cloudquery/test/versions/v1.2.3":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"name": "v1.2.3"}`))
			checkCreateAddonVersionRequest(t, r)
		case "/addons/cloudquery/test/versions/v1.2.3/assets":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`{"url": "%s"}`, "http://"+r.Host+"/upload-zip")))
		case "/upload-zip":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
		}
	}))
	defer ts.Close()

	cmd := NewCmdRoot()
	t.Setenv("CLOUDQUERY_API_URL", ts.URL)
	args := []string{"addon", "publish", "testdata/addon-v1/manifest.json", "v1.2.3"}
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(wantCalls, gotCalls); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func TestAddonPublishFinalize(t *testing.T) {
	t.Setenv("CLOUDQUERY_API_KEY", "testkey")

	wantCalls := map[string]int{
		"PUT /addons/cloudquery/test/versions/v1.2.3":         1,
		"POST /addons/cloudquery/test/versions/v1.2.3/assets": 1,
		"PUT /upload-zip": 1,
		"PATCH /addons/cloudquery/test/versions/v1.2.3": 1,
	}
	gotCalls := map[string]int{}
	gotUploads := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		gotCalls[r.Method+" "+r.URL.Path]++
		switch r.URL.Path {
		case "/addons/cloudquery/test/versions/v1.2.3":
			checkAuthHeader(t, r)
			if r.Method == "PATCH" {
				checkUpdateAddonVersionRequest(t, r)
				if gotUploads != 1 {
					t.Fatalf("expected 1 upload before draft=false, got %d", gotUploads)
				}
				w.WriteHeader(http.StatusOK)
			} else {
				checkCreateAddonVersionRequest(t, r)
				w.WriteHeader(http.StatusCreated)
			}
			w.Write([]byte(`{"name": "v1.2.3"}`))
		case "/addons/cloudquery/test/versions/v1.2.3/assets":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`{"url": "%s"}`, "http://"+r.Host+"/upload-zip")))
		case "/upload-zip":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
			gotUploads++
		}
	}))
	defer ts.Close()

	t.Setenv("CLOUDQUERY_API_URL", ts.URL)

	cmd := NewCmdRoot()
	args := []string{"addon", "publish", "testdata/addon-v1/manifest.json", "v1.2.3", "--finalize"}
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(wantCalls, gotCalls); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func TestAddonPublish_Unauthorized(t *testing.T) {
	t.Setenv("CLOUDQUERY_API_KEY", "badkey")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "unauthorized"}`))
	}))
	defer ts.Close()

	t.Setenv("CLOUDQUERY_API_URL", ts.URL)

	cmd := NewCmdRoot()
	args := []string{"addon", "publish", "testdata/addon-v1/manifest.json", "v1.2.3", "--finalize"}
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "unauthorized") {
		t.Fatalf("expected error to contain 'unauthorized', got %v", err)
	}
}

func checkCreateAddonVersionRequest(t *testing.T, r *http.Request) {
	t.Helper()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]any{}
	err = json.Unmarshal(body, &got)
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]any{
		"addon_deps":  nil,
		"checksum":    "b537240431bb4868264e48a8c646ebd3a9e355140d27d7fe559b5cbfd3ce6f31",
		"doc":         "# Test Addon README\n",
		"message":     "# Test Addon Changelog\n",
		"plugin_deps": []any{"cloudquery/source/test@v1.0.0"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func checkUpdateAddonVersionRequest(t *testing.T, r *http.Request) {
	t.Helper()

	got := map[string]any{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(body, &got)
	if err != nil {
		t.Fatal(err)
	}
	if got["draft"].(bool) {
		t.Fatalf("expected draft to be false, got %v", got["draft"])
	}
}
