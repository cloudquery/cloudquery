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

func createVirtualNetworkGatewayConnections(router *mux.Router) error {
	var connection armnetwork.VirtualNetworkGatewayConnectionsClientListResponse
	if err := faker.FakeObject(&connection); err != nil {
		return err
	}

	emptyStr := ""
	connection.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/connections", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&connection)
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

func TestVirtualNetworkGatewayConnections(t *testing.T) {
	client.MockTestHelper(t, VirtualNetworkGatewayConnections(), createVirtualNetworkGatewayConnections)
}
