package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPublish(t *testing.T) {
	t.Setenv("CQ_API_KEY", "testkey")

	wantCalls := map[string]int{
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3":                      1,
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3/tables":               1,
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3/docs":                 1,
		"POST /plugins/cloudquery/source/test/versions/v1.2.3/assets/linux_amd64":  1,
		"POST /plugins/cloudquery/source/test/versions/v1.2.3/assets/darwin_amd64": 1,
		"PUT /upload-linux":  1,
		"PUT /upload-darwin": 1,
	}
	gotCalls := map[string]int{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		gotCalls[r.Method+" "+r.URL.Path]++
		switch r.URL.Path {
		case "/plugins/cloudquery/source/test/versions/v1.2.3":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"name": "v1.2.3"}`))
			checkCreateVersionRequest(t, r)
		case "/plugins/cloudquery/source/test/versions/v1.2.3/tables":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{}`))
			checkCreateTablesRequest(t, r)
		case "/plugins/cloudquery/source/test/versions/v1.2.3/docs":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{}`))
			checkCreateDocsRequest(t, r)
		case "/plugins/cloudquery/source/test/versions/v1.2.3/assets/linux_amd64":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`{"url": "%s"}`, "http://"+r.Host+"/upload-linux")))
		case "/plugins/cloudquery/source/test/versions/v1.2.3/assets/darwin_amd64":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`{"url": "%s"}`, "http://"+r.Host+"/upload-darwin")))
		case "/upload-linux":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
		case "/upload-darwin":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
		}
	}))
	defer ts.Close()

	cmd := NewCmdRoot()
	args := []string{"publish", "cloudquery/test", "--dist-dir", "testdata/dist-v1", "--url", ts.URL}
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(wantCalls, gotCalls); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func TestPublishFinalize(t *testing.T) {
	t.Setenv("CQ_API_KEY", "testkey")

	wantCalls := map[string]int{
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3":                      1,
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3/tables":               1,
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3/docs":                 1,
		"POST /plugins/cloudquery/source/test/versions/v1.2.3/assets/linux_amd64":  1,
		"POST /plugins/cloudquery/source/test/versions/v1.2.3/assets/darwin_amd64": 1,
		"PUT /upload-linux":  1,
		"PUT /upload-darwin": 1,
		"PATCH /plugins/cloudquery/source/test/versions/v1.2.3": 1,
	}
	gotCalls := map[string]int{}
	gotUploads := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		gotCalls[r.Method+" "+r.URL.Path]++
		switch r.URL.Path {
		case "/plugins/cloudquery/source/test/versions/v1.2.3":
			checkAuthHeader(t, r)
			if r.Method == "PATCH" {
				checkUpdateVersionRequest(t, r)
				if gotUploads != 2 {
					t.Fatalf("expected 2 uploads before draft=false, got %d", gotUploads)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"name": "v1.2.3"}`))
			} else {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"name": "v1.2.3"}`))
				checkCreateVersionRequest(t, r)
			}
		case "/plugins/cloudquery/source/test/versions/v1.2.3/tables":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{}`))
			checkCreateTablesRequest(t, r)
		case "/plugins/cloudquery/source/test/versions/v1.2.3/docs":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{}`))
		case "/plugins/cloudquery/source/test/versions/v1.2.3/assets/linux_amd64":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`{"url": "%s"}`, "http://"+r.Host+"/upload-linux")))
		case "/plugins/cloudquery/source/test/versions/v1.2.3/assets/darwin_amd64":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(fmt.Sprintf(`{"url": "%s"}`, "http://"+r.Host+"/upload-darwin")))
		case "/upload-linux":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
			gotUploads++
		case "/upload-darwin":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
			gotUploads++
		}
	}))
	defer ts.Close()

	cmd := NewCmdRoot()
	args := []string{"publish", "cloudquery/test", "--dist-dir", "testdata/dist-v1", "--url", ts.URL, "--finalize"}
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(wantCalls, gotCalls); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func TestPublish_Unauthorized(t *testing.T) {
	t.Setenv("CQ_API_KEY", "badkey")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "unauthorized"}`))
	}))
	defer ts.Close()

	cmd := NewCmdRoot()
	args := []string{"publish", "cloudquery/test", "--dist-dir", "testdata/dist-v1", "--url", ts.URL, "--finalize"}
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "unauthorized") {
		t.Fatalf("expected error to contain 'unauthorized', got %v", err)
	}
}

func checkAuthHeader(t *testing.T, r *http.Request) {
	wantAuth := "Bearer testkey"
	if r.Header.Get("Authorization") != wantAuth {
		t.Fatalf("expected Authorization header to be %q, got %q", wantAuth, r.Header.Get("Authorization"))
	}
}

func checkCreateVersionRequest(t *testing.T, r *http.Request) {
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
		"message":      "## This is a subtitle\n\nThis is a paragraph.\n",
		"package_type": "native",
		"protocols":    []any{float64(3)},
		"supported_targets": []any{
			"linux_amd64", "darwin_amd64",
		},
		"checksums": []any{"1234567890", "1234567890"},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func checkUpdateVersionRequest(t *testing.T, r *http.Request) {
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

func checkCreateTablesRequest(t *testing.T, r *http.Request) {
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
		"tables": []any{
			map[string]any{
				"description":    "Test description",
				"is_incremental": false,
				"name":           "test_some_table",
				"relations":      []any{},
				"title":          "Test Some Table",
			},
		},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func checkCreateDocsRequest(t *testing.T, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]any{}
	err = json.Unmarshal(body, &got)
	if err != nil {
		t.Fatal(err)
	}

	customDocContent := readFile("testdata/dist-v1/docs/Custom-Doc.md")
	overviewContent := readFile("testdata/dist-v1/docs/overview.md")
	configurationContent := readFile("testdata/dist-v1/docs/configuration.md")

	want := map[string]any{
		"pages": []any{
			map[string]any{
				"content":          customDocContent,
				"name":             "custom-doc",
				"title":            "Custom Documentation",
				"ordinal_position": float64(3),
			},
			map[string]any{
				"content":          configurationContent,
				"name":             "configuration",
				"title":            "Configuration",
				"ordinal_position": float64(2),
			},
			map[string]any{
				"content":          overviewContent,
				"name":             "overview",
				"title":            "Overview",
				"ordinal_position": float64(1),
			},
		},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func readFile(name string) string {
	b, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return normalizeContent(string(b))
}
