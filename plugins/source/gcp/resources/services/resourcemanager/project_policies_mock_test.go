package resourcemanager

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	pb "google.golang.org/api/cloudresourcemanager/v3"
)

func createProjectPolicies(mux *httprouter.Router) error {
	var item pb.Policy
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	mux.POST("/v3/projects/testProject:getIamPolicy", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&item)
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

func TestProjectPolicies(t *testing.T) {
	client.MockTestRestHelper(t, ProjectPolicies(), createProjectPolicies, client.TestOptions{})
}
