package cloudsupport

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	pb "google.golang.org/api/cloudsupport/v2beta"
)

func createCases(mux *httprouter.Router) error {
	var item pb.ListCasesResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	item.NextPageToken = ""
	mux.GET("/v2beta/projects/testProject/cases", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func TestCases(t *testing.T) {
	client.MockTestRestHelper(t, Cases(), createCases, client.TestOptions{})
}
