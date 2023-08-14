package monitor

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createSubscriptionDiagnosticSettings(router *mux.Router) error {
	var item armmonitor.DiagnosticSettingsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	router.HandleFunc("/subscriptions/"+client.TestSubscription+"/providers/Microsoft.Insights/diagnosticSettings", func(w http.ResponseWriter, r *http.Request) {
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

func TestSubscriptionDiagnosticSettings(t *testing.T) {
	client.MockTestHelper(t, SubscriptionDiagnosticSettings(), createSubscriptionDiagnosticSettings)
}
