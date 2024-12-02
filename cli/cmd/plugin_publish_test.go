package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	gosync "sync"
	"testing"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/hub"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestPluginPublish(t *testing.T) {
	type testCase struct {
		name    string
		distDir string
		args    []string
	}

	var testCases = []testCase{
		{
			name:    "old package json with old command line args format",
			distDir: "testdata/dist-v1-no-team-package-json",
			args:    []string{"plugin", "publish", "cloudquery/test", "--dist-dir", "testdata/dist-v1-no-team-package-json"},
		},
		{
			name:    "new package json with old command line args format",
			distDir: "testdata/dist-v1-with-team-package-json",
			args:    []string{"plugin", "publish", "cloudquery/test", "--dist-dir", "testdata/dist-v1-with-team-package-json"},
		},
		{
			name:    "new package json with new command line args format (no args)",
			distDir: "testdata/dist-v1-with-team-package-json",
			args:    []string{"plugin", "publish", "--dist-dir", "testdata/dist-v1-with-team-package-json"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv("CLOUDQUERY_API_KEY", "testkey")
			wantCalls := map[string]int{
				"PUT /plugins/cloudquery/source/test/versions/v1.2.3":                      1,
				"PUT /plugins/cloudquery/source/test/versions/v1.2.3/tables":               1,
				"POST /plugins/cloudquery/source/test/versions/v1.2.3/docs":                1,
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
					checkCreatePluginVersionRequest(t, r)
				case "/plugins/cloudquery/source/test/versions/v1.2.3/tables":
					checkAuthHeader(t, r)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte(`{}`))
					checkCreateTablesRequest(t, r)
				case "/plugins/cloudquery/source/test/versions/v1.2.3/docs":
					checkAuthHeader(t, r)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte(`{}`))
					checkCreateDocsRequest(t, r, tc.distDir)
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
			t.Setenv(envAPIURL, ts.URL)
			allArgs := tc.args
			allArgs = append(allArgs, testCommandArgs(t)...)
			cmd.SetArgs(allArgs)
			err := cmd.Execute()
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(wantCalls, gotCalls); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestPluginPublishFinalize(t *testing.T) {
	t.Setenv("CLOUDQUERY_API_KEY", "testkey")

	wantCalls := map[string]int{
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3":                      1,
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3/tables":               1,
		"POST /plugins/cloudquery/source/test/versions/v1.2.3/docs":                1,
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
				checkUpdatePluginVersionRequest(t, r)
				if gotUploads != 2 {
					t.Fatalf("expected 2 uploads before draft=false, got %d", gotUploads)
				}
				w.WriteHeader(http.StatusOK)
			} else {
				checkCreatePluginVersionRequest(t, r)
				w.WriteHeader(http.StatusCreated)
			}
			w.Write([]byte(`{"name": "v1.2.3"}`))
		case "/plugins/cloudquery/source/test/versions/v1.2.3/tables":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{}`))
			checkCreateTablesRequest(t, r)
		case "/plugins/cloudquery/source/test/versions/v1.2.3/docs":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{}`))
			checkCreateDocsRequest(t, r, "testdata/dist-v1-with-team-package-json")
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

	t.Setenv(envAPIURL, ts.URL)

	cmd := NewCmdRoot()
	args := append([]string{"plugin", "publish", "--dist-dir", "testdata/dist-v1-with-team-package-json", "--finalize"}, testCommandArgs(t)...)
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(wantCalls, gotCalls); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func TestPluginPublishWithUI(t *testing.T) {
	const distDir = "testdata/dist-v1-with-team-package-json"

	t.Setenv("CLOUDQUERY_API_KEY", "testkey")
	wantCalls := map[string]int{
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3":                      1,
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3/tables":               1,
		"POST /plugins/cloudquery/source/test/versions/v1.2.3/docs":                1,
		"POST /plugins/cloudquery/source/test/versions/v1.2.3/assets/linux_amd64":  1,
		"POST /plugins/cloudquery/source/test/versions/v1.2.3/assets/darwin_amd64": 1,
		"PUT /upload-linux":   1,
		"PUT /upload-darwin":  1,
		"PUT /upload-uiasset": 2,
		"POST /plugins/cloudquery/source/test/versions/v1.2.3/uiassets": 1,
		"PUT /plugins/cloudquery/source/test/versions/v1.2.3/uiassets":  1,
	}
	gotCalls := map[string]int{}
	mu := &gosync.Mutex{}
	uiID := ""
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		mu.Lock()
		gotCalls[r.Method+" "+r.URL.Path]++
		mu.Unlock()
		// t.Log(r.Method, r.URL.Path)
		switch r.URL.Path {
		case "/plugins/cloudquery/source/test/versions/v1.2.3":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"name": "v1.2.3"}`))
			checkCreatePluginVersionRequest(t, r)
		case "/plugins/cloudquery/source/test/versions/v1.2.3/tables":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{}`))
			checkCreateTablesRequest(t, r)
		case "/plugins/cloudquery/source/test/versions/v1.2.3/docs":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{}`))
			checkCreateDocsRequest(t, r, distDir)
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
		case "/upload-uiasset":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{}`))
		case "/plugins/cloudquery/source/test/versions/v1.2.3/uiassets":
			checkAuthHeader(t, r)
			switch r.Method {
			case http.MethodPost:
				uiID = uuid.NewString()

				resp := cloudquery_api.UploadPluginUIAssets201Response{UIID: uiID}
				var rq cloudquery_api.UploadPluginUIAssetsRequest
				if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
					t.Fatal(err)
				}
				if len(rq.Assets) != 2 {
					t.Fatalf("expected 2 assets, got %d", len(rq.Assets))
				}
				for _, a := range rq.Assets {
					if a.Name != "index.html" && a.Name != "static/style.css" {
						t.Fatalf("unexpected asset name %q", a.Name)
					}
					resp.Assets = append(resp.Assets, cloudquery_api.PluginUIAsset{
						Name:      a.Name,
						UploadURL: "http://" + r.Host + "/upload-uiasset",
					})
				}

				w.WriteHeader(http.StatusCreated)
				if err := json.NewEncoder(w).Encode(resp); err != nil {
					t.Fatal(err)
				}
			case http.MethodPut:
				var rq cloudquery_api.FinalizePluginUIAssetUploadRequest
				if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
					t.Fatal(err)
				}
				if rq.UIID != uiID || rq.UIID == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					t.Fatalf("unexpected UIID %q", rq.UIID)
				}
				w.WriteHeader(http.StatusNoContent)
			default:
				w.WriteHeader(http.StatusNotAcceptable)
			}
		}
	}))
	defer ts.Close()

	cmd := NewCmdRoot()
	t.Setenv(envAPIURL, ts.URL)
	allArgs := []string{"plugin", "publish", "--dist-dir", "testdata/dist-v1-with-team-package-json", "--ui-dir", "testdata/ui-build"}
	allArgs = append(allArgs, testCommandArgs(t)...)
	cmd.SetArgs(allArgs)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(wantCalls, gotCalls); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func TestPluginPublish_Unauthorized(t *testing.T) {
	t.Setenv("CLOUDQUERY_API_KEY", "badkey")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "unauthorized"}`))
	}))
	defer ts.Close()

	t.Setenv(envAPIURL, ts.URL)

	cmd := NewCmdRoot()
	args := append([]string{"plugin", "publish", "--dist-dir", "testdata/dist-v1-with-team-package-json", "--finalize"}, testCommandArgs(t)...)
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
	t.Helper()

	wantAuth := "Bearer testkey"
	if r.Header.Get("Authorization") != wantAuth {
		t.Fatalf("expected Authorization header to be %q, got %q", wantAuth, r.Header.Get("Authorization"))
	}
}

func checkCreatePluginVersionRequest(t *testing.T, r *http.Request) {
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

func checkUpdatePluginVersionRequest(t *testing.T, r *http.Request) {
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

func checkCreateTablesRequest(t *testing.T, r *http.Request) {
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
		"tables": []any{
			map[string]any{
				"description":    "Test description",
				"is_incremental": false,
				"name":           "test_some_table",
				"relations":      []any{},
				"title":          "Test Some Table",
				"columns": []any{
					map[string]any{
						"name":            "_cq_id",
						"type":            "uuid",
						"description":     "Internal CQ ID of the row",
						"primary_key":     false,
						"not_null":        true,
						"unique":          true,
						"incremental_key": false,
					},
					map[string]any{
						"name":            "_cq_parent_id",
						"type":            "uuid",
						"description":     "Internal CQ ID of the parent row",
						"primary_key":     false,
						"not_null":        false,
						"unique":          false,
						"incremental_key": false,
					},
					map[string]any{
						"name":            "column1",
						"type":            "utf8",
						"description":     "Test Column 1",
						"primary_key":     true,
						"not_null":        false,
						"unique":          false,
						"incremental_key": false,
					},
					map[string]any{
						"name":            "column2",
						"type":            "int64",
						"description":     "Test Column 2",
						"primary_key":     false,
						"not_null":        false,
						"unique":          false,
						"incremental_key": false,
					},
					map[string]any{
						"name":            "client_id",
						"type":            "int64",
						"description":     "ID of client",
						"primary_key":     true,
						"not_null":        false,
						"unique":          false,
						"incremental_key": false,
					},
				},
			},
		},
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}

func checkCreateDocsRequest(t *testing.T, r *http.Request, distDir string) {
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

	customDocContent := readFile(distDir + "/docs/Custom-Doc.md")
	overviewContent := readFile(distDir + "/docs/overview.md")
	configurationContent := readFile(distDir + "/docs/configuration.md")

	want := map[string]any{
		"pages": []any{
			map[string]any{
				"content": customDocContent,
				"name":    "Custom-Doc",
			},
			map[string]any{
				"content": configurationContent,
				"name":    "configuration",
			},
			map[string]any{
				"content": overviewContent,
				"name":    "overview",
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
	return hub.NormalizeContent(string(b))
}
