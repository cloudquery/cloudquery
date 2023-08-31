package sql

import (
	"encoding/json"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createDatabases(router *mux.Router) error {
	var item armsql.DatabasesClientListByServerResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.Sql/servers/test string/databases", func(w http.ResponseWriter, r *http.Request) {
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
	if err := createDatabaseBlobAuditingPolicies(router); err != nil {
		return err
	}
	if err := createTransparentDataEncryptions(router); err != nil {
		return err
	}
	if err := createDatabaseVulnerabillityAssesments(router); err != nil {
		return err
	}
	if err := createLongTermRetentionPolicies(router); err != nil {
		return err
	}
	return createDatabaseThreatProtections(router)
}
