package authorization

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createClassicAdministrators(router *mux.Router) error {
	var item armauthorization.ClassicAdministratorsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/classicAdministrators", func(w http.ResponseWriter, r *http.Request) {
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

func TestClassicAdministrators(t *testing.T) {
	client.MockTestHelper(t, ClassicAdministrators(), createClassicAdministrators)
}
