package appservice

import (
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v2"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createVnetConnections(router *mux.Router) error {
	var item armappservice.WebAppsClientListVnetConnectionsResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.Web/sites/test string/virtualNetworkConnections", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item.VnetInfoResourceArray)
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
