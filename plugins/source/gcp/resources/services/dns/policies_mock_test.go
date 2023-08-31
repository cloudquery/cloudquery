package dns

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/dns/v1"
)

type MockPoliciesResult struct {
	Policies []*dns.Policy `json:"policies,omitempty"`
}

func createPolicies(mux *httprouter.Router) error {
	var item dns.Policy
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockPoliciesResult{
			Policies: []*dns.Policy{&item},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	return nil
}

func TestPolicies(t *testing.T) {
	client.MockTestRestHelper(t, Policies(), createPolicies, client.TestOptions{})
}
