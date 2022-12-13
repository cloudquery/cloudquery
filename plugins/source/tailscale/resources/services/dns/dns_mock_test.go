package dns

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func createDns(mux *httprouter.Router) error {
	var preferences tailscale.DNSPreferences
	if err := faker.FakeObject(&preferences); err != nil {
		return err
	}

	mux.GET("/api/v2/tailnet/:tailnet/dns/preferences", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(preferences)
		if err != nil {
			http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var nameservers []string
	if err := faker.FakeObject(&nameservers); err != nil {
		return err
	}

	mux.GET("/api/v2/tailnet/:tailnet/dns/nameservers", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(map[string][]string{
			"dns": nameservers,
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

func TestDns(t *testing.T) {
	client.MockTestHelper(t, Dns(), createDns)
}
