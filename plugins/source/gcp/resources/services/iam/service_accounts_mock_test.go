package iam

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/iam/v1"
)

func createServiceAccounts(mux *httprouter.Router) error {
	var item iam.ServiceAccount
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	item.Name = "test"
	mux.GET("/v1/projects/testProject/serviceAccounts", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &iam.ListServiceAccountsResponse{
			Accounts: []*iam.ServiceAccount{&item},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var key iam.ServiceAccountKey
	if err := faker.FakeObject(&key); err != nil {
		return err
	}
	key.ValidAfterTime = time.Now().Format(time.RFC3339)
	key.ValidBeforeTime = time.Now().Format(time.RFC3339)
	mux.GET("/v1/test/keys", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &iam.ListServiceAccountKeysResponse{
			Keys: []*iam.ServiceAccountKey{&key},
		}
		b, err := json.Marshal(resp)
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

func TestServiceAccounts(t *testing.T) {
	client.MockTestRestHelper(t, ServiceAccounts(), createServiceAccounts, client.TestOptions{})
}
