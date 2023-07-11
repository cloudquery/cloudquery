package organization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func createOrganizations(mux *httprouter.Router) error {
	var organization snyk.Organization
	if err := faker.FakeObject(&organization); err != nil {
		return err
	}
	organization.ID = "test-org-id"

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

	createOrganizationMembers(mux)
	createPendingProvisions(mux)
	return nil
}

func createOrganizationMembers(mux *httprouter.Router) {
	path := fmt.Sprintf("/org/%s/members", "test-org-id")
	mux.GET(path, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var member snyk.OrganizationMember
		if err := faker.FakeObject(&member); err != nil {
			http.Error(w, "unable to fake object: "+err.Error(), http.StatusBadRequest)
			return
		}
		b, err := json.Marshal([]snyk.OrganizationMember{member})
		if err != nil {
			http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
}

func createPendingProvisions(mux *httprouter.Router) {
	path := fmt.Sprintf("/org/%s/provision", "test-org-id")
	requestCount := 0
	mux.GET(path, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var pp snyk.PendingProvision
		defer func() { requestCount++ }()
		if err := faker.FakeObject(&pp); err != nil {
			http.Error(w, "unable to fake object: "+err.Error(), http.StatusBadRequest)
			return
		}
		var (
			b   []byte
			err error
		)
		if requestCount == 0 {
			b, err = json.Marshal([]snyk.PendingProvision{pp})
			if err != nil {
				http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			b, err = json.Marshal([]snyk.PendingProvision{})
			if err != nil {
				http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
				return
			}
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
}

func TestOrganizations(t *testing.T) {
	client.MockTestHelper(t, Organizations(), createOrganizations)
}
