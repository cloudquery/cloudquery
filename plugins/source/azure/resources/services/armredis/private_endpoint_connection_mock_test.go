// Code generated by codegen; DO NOT EDIT.

package armredis

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redis/armredis"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createPrivateEndpointConnection(router *mux.Router) error {
	var item armredis.PrivateEndpointConnectionsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/privateEndpointConnections", func(w http.ResponseWriter, r *http.Request) {
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

func TestPrivateEndpointConnection(t *testing.T) {
	client.MockTestHelper(t, PrivateEndpointConnection(), createPrivateEndpointConnection)
}
