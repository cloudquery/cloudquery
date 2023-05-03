package storage

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/gorilla/mux"
)

func createQueueAccessPolicy(router *mux.Router) error {
	var item azqueue.GetAccessPolicyResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/test string", func(w http.ResponseWriter, r *http.Request) {
		b, err := xml.Marshal(&item)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	return createAccountKey(router)
}

func createAccountKey(router *mux.Router) error {
	var keyResp armstorage.AccountsClientListKeysResponse
	if err := faker.FakeObject(&keyResp); err != nil {
		return err
	}
	keyResp.Keys[0].KeyName = to.Ptr("testkeyname")
	pp := armstorage.KeyPermissionRead
	keyResp.Keys[0].Permissions = &pp
	keyResp.Keys[0].Value = to.Ptr(base64.StdEncoding.EncodeToString([]byte("testkey")))

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.Storage/storageAccounts/test string/listKeys", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&keyResp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println(string(b))
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}
