package storage

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createAccounts(router *mux.Router) error {
	var item armstorage.AccountsClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	val1 := *item.Value[0]
	val1.ID = to.Ptr("testaccount1")
	val1.Name = to.Ptr("testaccount1")

	val2 := *item.Value[0]
	val2.ID = to.Ptr("testaccount2")
	val2.Name = to.Ptr("testaccount2")

	val3 := *item.Value[0]
	val3.ID = to.Ptr("testaccount3")
	val3.Name = to.Ptr("testaccount3")

	item.Value = []*armstorage.Account{&val1, &val2, &val3}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Storage/storageAccounts", func(w http.ResponseWriter, r *http.Request) {
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
	if err := createTables(router); err != nil {
		return err
	}
	if err := createContainers(router); err != nil {
		return err
	}
	if err := createFileShares(router); err != nil {
		return err
	}
	if err := createBlobServices(router); err != nil {
		return err
	}
	if err := createQueueServices(router); err != nil {
		return err
	}
	return createQueues(router)
}

func TestAccounts(t *testing.T) {
	client.MockTestHelper(t, Accounts(), createAccounts)
}
