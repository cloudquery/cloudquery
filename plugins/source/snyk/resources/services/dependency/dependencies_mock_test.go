package dependency

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func createDependencies(mux *httprouter.Router) error {
	var dependency snyk.Dependency
	if err := faker.FakeObject(&dependency); err != nil {
		return err
	}

	mux.POST("/org/:orgID/dependencies", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		type DependenciesRoot struct {
			Total   int               `json:"total,omitempty"`
			Results []snyk.Dependency `json:"results,omitempty"`
		}
		b, err := json.Marshal(DependenciesRoot{
			Total:   1,
			Results: []snyk.Dependency{dependency},
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

func TestDependencies(t *testing.T) {
	client.MockTestHelper(t, Dependencies(), createDependencies)
}
