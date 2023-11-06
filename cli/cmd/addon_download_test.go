package cmd

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestAddonDownload(t *testing.T) {
	t.Setenv("CLOUDQUERY_API_KEY", "testkey")

	tempDir := t.TempDir()
	expectedFileAndPath := filepath.Join(tempDir, "cloudquery_visualization_test_v1.2.3.zip")

	wantCalls := map[string]int{
		"GET /addons/cloudquery/visualization/test":                        1,
		"GET /addons/cloudquery/visualization/test/versions/v1.2.3":        1,
		"GET /addons/cloudquery/visualization/test/versions/v1.2.3/assets": 1,
		"GET /asset-zip": 1,
	}

	payload := []byte("test payload")

	s := sha256.New()
	s.Write(payload)
	payloadChecksum := fmt.Sprintf("%x", s.Sum(nil))

	gotCalls := map[string]int{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s %s", r.Method, r.URL.Path)
			t.FailNow()
		}
		w.Header().Set("Content-Type", "application/json")
		gotCalls[r.Method+" "+r.URL.Path]++
		switch r.URL.Path {
		case "/addons/cloudquery/visualization/test":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusOK)
			b, _ := json.Marshal(cloudquery_api.Addon{
				AddonFormat: cloudquery_api.Zip,
				AddonType:   cloudquery_api.Visualization,
				Name:        "test",
				TeamName:    "cloudquery",
			})
			w.Write(b)
		case "/addons/cloudquery/visualization/test/versions/v1.2.3":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusOK)
			b, _ := json.Marshal(cloudquery_api.AddonVersion{
				Checksum: payloadChecksum,
				Name:     "v1.2.3",
			})
			w.Write(b)
		case "/addons/cloudquery/visualization/test/versions/v1.2.3/assets":
			checkAuthHeader(t, r)
			w.Header().Set("Location", "http://"+r.Host+"/asset-zip")
			w.WriteHeader(http.StatusFound)
		case "/asset-zip":
			w.Header().Set("Content-Type", "application/octet-stream")
			w.WriteHeader(http.StatusOK)
			w.Write(payload)
		}
	}))
	defer ts.Close()

	cmd := NewCmdRoot()
	t.Setenv(envAPIURL, ts.URL)
	args := []string{"addon", "download", "cloudquery/visualization/test@v1.2.3", "-t", tempDir}
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(wantCalls, gotCalls); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}

	_, err = os.Stat(expectedFileAndPath)
	assert.NoError(t, err, "expected file %s to exist", expectedFileAndPath)
}

func TestAddonDownloadStdout(t *testing.T) {
	t.Setenv("CLOUDQUERY_API_KEY", "testkey")

	wantCalls := map[string]int{
		"GET /addons/cloudquery/visualization/test":                        1,
		"GET /addons/cloudquery/visualization/test/versions/v1.2.3":        1,
		"GET /addons/cloudquery/visualization/test/versions/v1.2.3/assets": 1,
		"GET /asset-zip": 1,
	}

	payload := []byte("payload to stdout")

	s := sha256.New()
	s.Write(payload)
	payloadChecksum := fmt.Sprintf("%x", s.Sum(nil))

	gotCalls := map[string]int{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s %s", r.Method, r.URL.Path)
			t.FailNow()
		}
		w.Header().Set("Content-Type", "application/json")
		gotCalls[r.Method+" "+r.URL.Path]++
		switch r.URL.Path {
		case "/addons/cloudquery/visualization/test":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusOK)
			b, _ := json.Marshal(cloudquery_api.Addon{
				AddonFormat: cloudquery_api.Zip,
				AddonType:   cloudquery_api.Visualization,
				Name:        "test",
				TeamName:    "cloudquery",
			})
			w.Write(b)
		case "/addons/cloudquery/visualization/test/versions/v1.2.3":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusOK)
			b, _ := json.Marshal(cloudquery_api.AddonVersion{
				Checksum: payloadChecksum,
				Name:     "v1.2.3",
			})
			w.Write(b)
		case "/addons/cloudquery/visualization/test/versions/v1.2.3/assets":
			checkAuthHeader(t, r)
			w.Header().Set("Location", "http://"+r.Host+"/asset-zip")
			w.WriteHeader(http.StatusFound)
		case "/asset-zip":
			w.Header().Set("Content-Type", "application/octet-stream")
			w.WriteHeader(http.StatusOK)
			w.Write(payload)
		}
	}))
	defer ts.Close()

	cmd := NewCmdRoot()
	t.Setenv(envAPIURL, ts.URL)
	args := []string{"addon", "download", "cloudquery/visualization/test@v1.2.3", "-t", "-"}
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(wantCalls, gotCalls); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
}
