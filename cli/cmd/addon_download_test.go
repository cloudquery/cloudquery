package cmd

import (
	"crypto/sha256"
	"encoding/json"
	"io"

	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddonDownload(t *testing.T) {
	t.Setenv("CLOUDQUERY_API_KEY", "testkey")

	tempDir := t.TempDir()
	expectedFileAndPath := filepath.Join(tempDir, "cloudquery_visualization_test_v1.2.3.zip")

	wantCalls := map[string]int{
		"GET /teams": 1,
		"GET /teams/test_team/addons/cloudquery/visualization/test/versions/v1.2.3/assets":                1,
		"GET /assets/cloudquery/addon_visualization/test/v1.2.3/cloudquery_visualization_test_v1.2.3.zip": 1,
	}

	payload := []byte("test payload")

	s := sha256.New()
	s.Write(payload)
	payloadChecksum := hex.EncodeToString(s.Sum(nil))

	gotCalls := map[string]int{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		gotCalls[r.Method+" "+r.URL.Path]++
		switch r.URL.Path {
		case "/teams/test_team/addons/cloudquery/visualization/test/versions/v1.2.3/assets":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusOK)
			b, _ := json.Marshal(map[string]string{
				"checksum": payloadChecksum,
				"location": "http://" + r.Host + "/assets/cloudquery/addon_visualization/test/v1.2.3/cloudquery_visualization_test_v1.2.3.zip",
			})
			_, err := w.Write(b)
			require.NoError(t, err)
		case "/assets/cloudquery/addon_visualization/test/v1.2.3/cloudquery_visualization_test_v1.2.3.zip":
			w.Header().Set("Content-Type", "application/octet-stream")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(payload)
			require.NoError(t, err)
		case "/teams":
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"items":[{"name":"test_team","displayName":"Test Team"}]}`))
			require.NoError(t, err)
		}
	}))
	defer ts.Close()

	cmd := NewCmdRoot()
	t.Setenv(envAPIURL, ts.URL)
	args := append([]string{"addon", "download", "cloudquery/visualization/test@v1.2.3", "-t", tempDir}, testCommandArgs(t)...)
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
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	t.Setenv("CLOUDQUERY_API_KEY", "testkey")

	wantCalls := map[string]int{
		"GET /teams": 1,
		"GET /teams/test_team/addons/cloudquery/visualization/test/versions/v1.2.3/assets":                1,
		"GET /assets/cloudquery/addon_visualization/test/v1.2.3/cloudquery_visualization_test_v1.2.3.zip": 1,
	}

	payload := []byte("payload to stdout")

	s := sha256.New()
	s.Write(payload)
	payloadChecksum := hex.EncodeToString(s.Sum(nil))

	gotCalls := map[string]int{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		gotCalls[r.Method+" "+r.URL.Path]++
		switch r.URL.Path {
		case "/teams/test_team/addons/cloudquery/visualization/test/versions/v1.2.3/assets":
			checkAuthHeader(t, r)
			w.WriteHeader(http.StatusOK)
			b, _ := json.Marshal(map[string]string{
				"checksum": payloadChecksum,
				"location": "http://" + r.Host + "/assets/cloudquery/addon_visualization/test/v1.2.3/cloudquery_visualization_test_v1.2.3.zip",
			})
			_, err := w.Write(b)
			require.NoError(t, err)
		case "/assets/cloudquery/addon_visualization/test/v1.2.3/cloudquery_visualization_test_v1.2.3.zip":
			w.Header().Set("Content-Type", "application/octet-stream")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(payload)
			require.NoError(t, err)
		case "/teams":
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`{"items":[{"name":"test_team","displayName":"Test Team"}]}`))
			require.NoError(t, err)
		}
	}))
	defer ts.Close()

	cmd := NewCmdRoot()
	t.Setenv(envAPIURL, ts.URL)
	args := append([]string{"addon", "download", "cloudquery/visualization/test@v1.2.3", "-t", "-"}, testCommandArgs(t)...)
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(wantCalls, gotCalls); diff != "" {
		t.Fatalf("mismatch (-want +got):\n%s", diff)
	}
	out, _ := io.ReadAll(r)
	require.Equal(t, payload, out)
}
