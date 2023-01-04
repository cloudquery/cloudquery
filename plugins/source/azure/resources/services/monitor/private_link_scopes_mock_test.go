// Code generated by codegen; DO NOT EDIT.
package monitor

import (
	"encoding/json"
	"net/http"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createPrivateLinkScopes(router *mux.Router) error {
	var item armmonitor.PrivateLinkScopesClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/microsoft.insights/privateLinkScopes", func(w http.ResponseWriter, r *http.Request) {
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

func TestPrivateLinkScopes(t *testing.T) {
	client.MockTestHelper(t, PrivateLinkScopes(), createPrivateLinkScopes)
}
