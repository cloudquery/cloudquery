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

func createAutoscaleSettings(router *mux.Router) error {
	var item armmonitor.AutoscaleSettingsClientListBySubscriptionResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyString := ""
	item.NextLink = &emptyString

	router.HandleFunc("/subscriptions/"+client.TestSubscription+"/providers/Microsoft.Insights/autoscalesettings", func(w http.ResponseWriter, r *http.Request) {
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

func TestAutoscaleSettings(t *testing.T) {
	client.MockTestHelper(t, AutoscaleSettings(), createAutoscaleSettings)
}
