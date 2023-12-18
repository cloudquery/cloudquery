package users

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/fakes"
	"github.com/gorilla/mux"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func handleUsers(router *mux.Router) error {
	u := fakes.User()
	u.AdditionalProperties = map[string]any{"key": "value"}
	u.Embedded = map[string]map[string]any{"top-key": {"key": "value"}}
	u.Links = map[string]map[string]any{"top-key": {"key": "value"}}

	router.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		b, err := json.Marshal([]okta.User{u})
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

func TestUsers(t *testing.T) {
	client.MockTestHelper(t, Users(), handleUsers)
}
