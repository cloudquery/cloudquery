// Code generated by codegen; DO NOT EDIT.

package armapimanagement

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createUserContract(router *mux.Router) error {
	var item armapimanagement.GroupUserClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/groups/{groupId}/users", func(w http.ResponseWriter, r *http.Request) {
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

func TestUserContract(t *testing.T) {
	client.MockTestHelper(t, UserContract(), createUserContract)
}
