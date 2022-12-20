package groups

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/okta/client"
	"github.com/cloudquery/cloudquery/plugins/source/okta/resources/services/fakes"
	"github.com/gorilla/mux"
	"github.com/okta/okta-sdk-golang/v3/okta"
)

func handleGroups(router *mux.Router) error {
	g := fakes.Group()

	router.HandleFunc("/api/v1/groups", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		b, err := json.Marshal([]okta.Group{g})
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	router.HandleFunc("/api/v1/groups/"+*g.Id+"/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		b, err := json.Marshal([]okta.User{fakes.User()})
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

func TestGroups(t *testing.T) {
	client.MockTestHelper(t, Groups(), handleGroups)
}
