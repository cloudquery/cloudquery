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

func createGroups(router *mux.Router) error {
	var groupInfo armmanagementgroups.ManagementGroupListResult
	if err := faker.FakeObject(&groupInfo); err != nil {
		return err
	}

	router.HandleFunc("/providers/Microsoft.Management/managementGroups", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&groupInfo)
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

func TestManagementGroups(t *testing.T) {
	client.MockTestHelper(t, ManagementGroups(), createGroups)
}
