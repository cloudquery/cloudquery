package storage

import (
	"encoding/xml"
	"net/http"

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
	return nil
}
