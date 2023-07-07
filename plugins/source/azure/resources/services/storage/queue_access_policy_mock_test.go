package storage

import (
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azqueue"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createQueueAccessPolicy(router *mux.Router) error {
	var item azqueue.GetAccessPolicyResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	for _, qName := range []string{"acc1testqueue1", "acc1testqueue2", "acc2testqueue1", "acc2testqueue2", "acc3testqueue1", "acc3testqueue2"} {
		qName := qName
		router.HandleFunc("/"+qName, func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if !strings.HasPrefix(auth, "SharedKey ") {
				http.Error(w, "SharedKey auth is required", http.StatusUnauthorized)
				return
			}
			keyAcc := strings.Split(strings.TrimPrefix(auth, "SharedKey "), ":")[0]
			var wantAcc string
			switch {
			case strings.HasPrefix(qName, "acc1"):
				wantAcc = "testaccount1"
			case strings.HasPrefix(qName, "acc2"):
				wantAcc = "testaccount2"
			case strings.HasPrefix(qName, "acc3"):
				wantAcc = "testaccount3"
			}
			if keyAcc != wantAcc {
				http.Error(w, "invalid account got: "+keyAcc+" want: "+wantAcc, http.StatusUnauthorized)
				return
			}

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
	}
	return createAccountKey(router)
}

func createAccountKey(router *mux.Router) error {
	var keyResp armstorage.AccountsClientListKeysResponse
	if err := faker.FakeObject(&keyResp); err != nil {
		return err
	}
	keyResp.Keys[0].KeyName = to.Ptr("testkeyname")
	keyResp.Keys[0].Permissions = to.Ptr(armstorage.KeyPermissionRead)
	keyResp.Keys[0].Value = to.Ptr(base64.StdEncoding.EncodeToString([]byte("testkey")))

	mu := &sync.Mutex{}
	reqCounts := make(map[string]int64)

	router.HandleFunc("/subscriptions/{subscriptionId}/resourceGroups/debug/providers/Microsoft.Storage/storageAccounts/{accountName}/listKeys", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		if reqCounts[r.RequestURI] > 0 {
			http.Error(w, "called more than once: "+r.RequestURI, http.StatusInternalServerError)
			return
		}

		reqCounts[r.RequestURI]++

		var resp armstorage.AccountsClientListKeysResponse
		if !strings.Contains(r.RequestURI, "testaccount3") {
			resp = keyResp
		}

		b, err := json.Marshal(&resp)
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
