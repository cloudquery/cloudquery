package iam

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/iam/v1"
)

type MockRolesResult struct {
	Roles []*iam.Role `json:"roles,omitempty"`
}

func createRoles(mux *httprouter.Router) error {
	var item iam.Role
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockRolesResult{
			Roles: []*iam.Role{&item},
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

func TestRoles(t *testing.T) {
	client.MockTestRestHelper(t, Roles(), createRoles, client.TestOptions{})
}
