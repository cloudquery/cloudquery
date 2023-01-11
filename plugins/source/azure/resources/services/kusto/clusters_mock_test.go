// Code generated by codegen; DO NOT EDIT.
package kusto

import (
	"encoding/json"
	"net/http"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createClusters(router *mux.Router) error {
	var item armkusto.ClustersClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/clusters", func(w http.ResponseWriter, r *http.Request) {
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

func TestClusters(t *testing.T) {
	client.MockTestHelper(t, Clusters(), createClusters)
}
