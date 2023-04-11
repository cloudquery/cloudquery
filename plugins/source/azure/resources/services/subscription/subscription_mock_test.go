package subscription

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/gorilla/mux"
)

func createSubscriptions(router *mux.Router) error {
	var item armsubscription.SubscriptionsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	emptyStr := ""
	item.NextLink = &emptyStr
	item.Value[0].ID = to.Ptr("/subscriptions/sub-id")
	item.Value[0].SubscriptionID = to.Ptr("sub-id")

	router.HandleFunc("/subscriptions", func(w http.ResponseWriter, r *http.Request) {
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

	return createLocations(router)
}

func TestSubscriptions(t *testing.T) {
	client.MockTestHelper(t, Subscriptions(), createSubscriptions)
}
