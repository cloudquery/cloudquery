package dns

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/tailscale/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
)

func createNameservers(mux *httprouter.Router) error {
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

	return nil
}

func TestNameservers(t *testing.T) {
	client.MockTestHelper(t, Nameservers(), createNameservers)
}
