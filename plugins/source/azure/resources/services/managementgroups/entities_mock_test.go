package managementgroups

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createEntities(router *mux.Router) error {
	var entityInfo armmanagementgroups.EntitiesClientListResponse
	if err := faker.FakeObject(&entityInfo); err != nil {
		return err
	}

	router.HandleFunc("/providers/Microsoft.Management/getEntities", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&entityInfo)
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

func TestEntities(t *testing.T) {
	client.MockTestHelper(t, Entities(), createEntities)
}
