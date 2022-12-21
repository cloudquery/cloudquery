package iam

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	iam "google.golang.org/api/iam/v2beta"
)

func createDenyPolicies(mux *httprouter.Router) error {
	var item iam.GoogleIamV2betaPolicy
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &iam.GoogleIamV2betaListPoliciesResponse{
			Policies: []*iam.GoogleIamV2betaPolicy{&item},
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

func TestDenyPolicies(t *testing.T) {
	client.MockTestRestHelper(t, DenyPolicies(), createDenyPolicies, client.TestOptions{})
}
