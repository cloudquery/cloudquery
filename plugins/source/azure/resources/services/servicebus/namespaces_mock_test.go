package servicebus

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createNamespaces(router *mux.Router) error {
	var item armservicebus.NamespacesClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.ServiceBus/namespaces", func(w http.ResponseWriter, r *http.Request) {
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

	err := createTopics(router)
	if err != nil {
		return err
	}

	err = createTopicAuthorizationRules(router)
	if err != nil {
		return err
	}

	return createTopicRuleAccessKeys(router)
}

func TestNamespaces(t *testing.T) {
	client.MockTestHelper(t, Namespaces(), createNamespaces)
}
