package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPublish(t *testing.T) {
	gotRequests := make([]*http.Request, 0)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotRequests = append(gotRequests, r)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"success": true}`))
	}))
	defer ts.Close()

	cmd := NewCmdRoot()
	args := []string{"publish", "cloudquery/test", "--dist", "testdata/dist-v1", "--url", ts.URL}
	cmd.SetArgs(args)
	err := cmd.Execute()
	if err != nil {
		t.Fatal(err)
	}
	if len(gotRequests) != 1 {
		t.Fatalf("unexpected number of requests: %d", len(gotRequests))
	}
}
