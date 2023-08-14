package azurearcdata

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createPostgresInstances(router *mux.Router) error {
	var item armazurearcdata.PostgresInstancesClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.AzureArcData/postgresInstances", func(w http.ResponseWriter, r *http.Request) {
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

func TestPostgresInstances(t *testing.T) {
	client.MockTestHelper(t, PostgresInstances(), createPostgresInstances)
}
