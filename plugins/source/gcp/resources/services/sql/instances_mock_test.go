package sql

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

func createInstances(mux *httprouter.Router) error {
	var item sql.InstancesListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	item.NextPageToken = ""

	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(item)
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

func TestInstances(t *testing.T) {
	client.MockTestRestHelper(t, Instances(), createInstances, client.TestOptions{})
}
