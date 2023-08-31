package group

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

func createGroups(mux *httprouter.Router) error {
	return createGroupMembers(mux, "test-group-id")
}

func createGroupMembers(mux *httprouter.Router, groupID string) error {
	path := fmt.Sprintf("/group/%s/members", groupID)
	mux.GET(path, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var member snyk.GroupMember
		if err := faker.FakeObject(&member); err != nil {
			http.Error(w, "unable to fake object: "+err.Error(), http.StatusBadRequest)
			return
		}
		b, err := json.Marshal([]snyk.GroupMember{member})
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
	client.MockTestHelper(t, Groups(), createGroups)
}
