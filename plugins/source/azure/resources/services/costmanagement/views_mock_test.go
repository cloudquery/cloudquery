package costmanagement

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/gorilla/mux"
)

func createViews(router *mux.Router) error {
	var item armcostmanagement.ViewsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.CostManagement/views", func(w http.ResponseWriter, r *http.Request) {
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
	return createViewQueries(router)
}

func TestViews(t *testing.T) {
	client.MockTestHelper(t, Views(), createViews)
}
