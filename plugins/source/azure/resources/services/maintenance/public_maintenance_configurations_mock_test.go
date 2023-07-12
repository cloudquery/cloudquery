package maintenance

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createPublicMaintenanceConfigurations(router *mux.Router) error {
	var item armmaintenance.PublicMaintenanceConfigurationsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Maintenance/publicMaintenanceConfigurations", func(w http.ResponseWriter, r *http.Request) {
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

func TestPublicMaintenanceConfigurations(t *testing.T) {
	client.MockTestHelper(t, PublicMaintenanceConfigurations(), createPublicMaintenanceConfigurations)
}
