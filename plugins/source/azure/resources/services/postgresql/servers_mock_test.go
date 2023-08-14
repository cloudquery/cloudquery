package postgresql

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createServers(router *mux.Router) error {
	var item armpostgresql.ServersClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.DBforPostgreSQL/servers", func(w http.ResponseWriter, r *http.Request) {
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

	if err := createConfigurations(router); err != nil {
		return err
	}
	if err := createFirewallRules(router); err != nil {
		return err
	}
	return createDatabases(router)
}

func TestServers(t *testing.T) {
	client.MockTestHelper(t, Servers(), createServers)
}
