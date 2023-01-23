package organization

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func createOrganizations(mux *httprouter.Router) error {
	var organization snyk.Organization
	if err := faker.FakeObject(&organization); err != nil {
		return err
	}
	organization.ID = "test-org"

	mux.GET("/orgs", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		type organizationsRoot struct {
			Organizations []snyk.Organization `json:"orgs,omitempty"`
		}
		b, err := json.Marshal(organizationsRoot{Organizations: []snyk.Organization{organization}})
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

func TestOrganizations(t *testing.T) {
	client.MockTestHelper(t, Organizations(), createOrganizations)
}
