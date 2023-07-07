package network

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createVirtualNetworkGateways(router *mux.Router) error {
	var gateway armnetwork.VirtualNetworkGatewaysClientListResponse
	if err := faker.FakeObject(&gateway); err != nil {
		return err
	}

	emptyStr := ""
	gateway.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworkGateways", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&gateway)
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

func TestVirtualNetworkGateways(t *testing.T) {
	client.MockTestHelper(t, VirtualNetworkGateways(), createVirtualNetworkGateways)
}
