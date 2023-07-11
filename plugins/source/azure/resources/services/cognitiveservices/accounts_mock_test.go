package cognitiveservices

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createAccounts(router *mux.Router) error {
	var item armcognitiveservices.AccountsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/accounts", func(w http.ResponseWriter, r *http.Request) {
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

	if err := createAccountDeployments(router); err != nil {
		return err
	}

	if err := createAccountPrivateEndpointConnections(router); err != nil {
		return err
	}

	if err := createAccountPrivateLinkResources(router); err != nil {
		return err
	}

	if err := createAccountModels(router); err != nil {
		return err
	}

	if err := createAccountUsages(router); err != nil {
		return err
	}

	return createAccountSKUs(router)
}

func TestAccounts(t *testing.T) {
	client.MockTestHelper(t, Accounts(), createAccounts)
}
