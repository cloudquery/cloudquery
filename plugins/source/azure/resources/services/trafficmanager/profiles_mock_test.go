package trafficmanager

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createProfiles(router *mux.Router) error {
	var item armtrafficmanager.ProfilesClientListBySubscriptionResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Network/trafficmanagerprofiles",
		func(w http.ResponseWriter, r *http.Request) {
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

func TestProfiles(t *testing.T) {
	client.MockTestHelper(t, Profiles(), createProfiles)
}
