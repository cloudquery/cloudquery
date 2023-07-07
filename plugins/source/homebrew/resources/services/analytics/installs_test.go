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

func buildInstalls(t *testing.T) *homebrew.Client {
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

func TestInstalls30d(t *testing.T) {
	client.MockTestHelper(t, Installs(homebrew.Days30), buildInstalls, client.TestOptions{})
}

func TestInstalls90d(t *testing.T) {
	client.MockTestHelper(t, Installs(homebrew.Days90), buildInstalls, client.TestOptions{})
}

func TestInstalls365d(t *testing.T) {
	client.MockTestHelper(t, Installs(homebrew.Days365), buildInstalls, client.TestOptions{})
}
