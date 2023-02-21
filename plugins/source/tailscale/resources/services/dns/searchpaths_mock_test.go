package dns

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
)

func createSearchpaths(mux *httprouter.Router) error {
	var searchpaths []string
	if err := faker.FakeObject(&searchpaths); err != nil {
		return err
	}
	mux.GET("/api/v2/tailnet/:tailnet/dns/searchpaths", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(map[string][]string{
			"searchPaths": searchpaths,
		})
		if err != nil {
			http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestSearchpaths(t *testing.T) {
	client.MockTestHelper(t, Searchpaths(), createSearchpaths)
}
