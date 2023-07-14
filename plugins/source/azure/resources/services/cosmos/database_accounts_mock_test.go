package cosmos

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createDatabaseAccounts(router *mux.Router) error {
	var item armcosmos.DatabaseAccountsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/databaseAccounts", func(w http.ResponseWriter, r *http.Request) {
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
	if err := createMongoDbDatabases(router); err != nil {
		return err
	}
	return createSqlDatabases(router)
}

func TestDatabaseAccounts(t *testing.T) {
	client.MockTestHelper(t, DatabaseAccounts(), createDatabaseAccounts)
}
