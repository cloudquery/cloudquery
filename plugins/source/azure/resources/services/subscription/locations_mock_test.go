package subscription

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
	"net/http"
)

func createLocations(router *mux.Router) error {
	var item armsubscription.SubscriptionsClientListLocationsResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/test string/locations", func(w http.ResponseWriter, r *http.Request) {
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
