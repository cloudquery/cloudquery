package analytics

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/homebrew/client"
	"github.com/cloudquery/cloudquery/plugins/source/homebrew/internal/homebrew"
	"github.com/cloudquery/plugin-sdk/v4/faker"
)

func buildErrors(t *testing.T) *homebrew.Client {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		installs := homebrew.InstallsResponse{}
		err := faker.FakeObject(&installs)
		if err != nil {
			t.Fatal(err)
		}
		j, err := json.Marshal(installs)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(j)
	}))
	t.Cleanup(ts.Close)
	c := homebrew.NewClient(
		homebrew.WithBaseURL(ts.URL),
		homebrew.WithHTTPClient(ts.Client()),
	)
	return c
}

func TestBuildErrors30d(t *testing.T) {
	client.MockTestHelper(t, BuildErrors(homebrew.Days30), buildErrors, client.TestOptions{})
}

func TestBuildErrors90d(t *testing.T) {
	client.MockTestHelper(t, BuildErrors(homebrew.Days90), buildErrors, client.TestOptions{})
}

func TestBuildErrors365d(t *testing.T) {
	client.MockTestHelper(t, BuildErrors(homebrew.Days365), buildErrors, client.TestOptions{})
}
