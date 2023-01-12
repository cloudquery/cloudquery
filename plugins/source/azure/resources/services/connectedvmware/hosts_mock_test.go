package connectedvmware

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createHosts(router *mux.Router) error {
	var item armconnectedvmware.HostsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.ConnectedVMwarevSphere/hosts", func(w http.ResponseWriter, r *http.Request) {
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

func TestHosts(t *testing.T) {
	client.MockTestHelper(t, Hosts(), createHosts)
}
