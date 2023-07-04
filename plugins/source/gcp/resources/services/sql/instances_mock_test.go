package sql

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

func createInstances(mux *httprouter.Router) error {
	var instanceResponse sql.InstancesListResponse
	if err := faker.FakeObject(&instanceResponse); err != nil {
		return err
	}
	instanceResponse.NextPageToken = ""

	mux.GET("/sql/v1beta4/projects/testProject/instances", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(instanceResponse)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var usersResponse sql.UsersListResponse
	if err := faker.FakeObject(&usersResponse); err != nil {
		return err
	}
	usersResponse.NextPageToken = ""

	mux.GET("/sql/v1beta4/projects/testProject/instances/test string/users", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&usersResponse)
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
