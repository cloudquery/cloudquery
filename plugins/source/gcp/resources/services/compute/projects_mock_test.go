package compute

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	pb "cloud.google.com/go/compute/apiv1/computepb"
)

func createProjects(mux *httprouter.Router) error {
	var item pb.Project
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func TestProjects(t *testing.T) {
	client.MockTestRestHelper(t, Projects(), createProjects, client.TestOptions{})
}
