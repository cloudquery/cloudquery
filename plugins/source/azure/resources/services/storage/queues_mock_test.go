package storage

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createQueues(router *mux.Router) error {
	// Create two queues for each account, and prefix them with "acc1" or "acc2"
	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/default/queues", func(w http.ResponseWriter, r *http.Request) {
		var item armstorage.QueueClientListResponse
		if err := faker.FakeObject(&item); err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		tempVal := *item.Value[0]
		item.NextLink = to.Ptr("")

		prefix := ""
		switch {
		case strings.Contains(r.RequestURI, "/storageAccounts/testaccount1/"):
			prefix = "acc1"
		case strings.Contains(r.RequestURI, "/storageAccounts/testaccount2/"):
			prefix = "acc2"
		case strings.Contains(r.RequestURI, "/storageAccounts/testaccount3/"):
			prefix = "acc3"
		}

		val1 := tempVal
		val1.ID = to.Ptr(prefix + "testqueue1")
		val1.Name = to.Ptr(prefix + "testqueue1")

		val2 := tempVal
		val2.ID = to.Ptr(prefix + "testqueue2")
		val2.Name = to.Ptr(prefix + "testqueue2")

		item.Value = []*armstorage.ListQueue{&val1, &val2}

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
	return createQueueAccessPolicy(router)
}
