package monitor

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/gorilla/mux"
)

func createLogProfiles(router *mux.Router) error {
	var item armmonitor.LogProfilesClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/logprofiles", func(w http.ResponseWriter, r *http.Request) {
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

func TestLogProfiles(t *testing.T) {
	client.MockTestHelper(t, LogProfiles(), createLogProfiles)
}
