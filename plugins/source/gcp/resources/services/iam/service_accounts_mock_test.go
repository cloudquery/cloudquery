package iam

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/iam/v1"

	"google.golang.org/api/option"
)

func createServiceAccounts() (*client.Services, error) {
	var item iam.ServiceAccount
	if err := faker.FakeObject(&item); err != nil {
		return nil, err
	}
	item.Name = "test"
	mux := httprouter.New()
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
		return nil, err
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

	ts := httptest.NewServer(mux)
	svc, err := iam.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Iam: svc,
	}, nil
}

func TestServiceAccounts(t *testing.T) {
	client.MockTestHelper(t, ServiceAccounts(), createServiceAccounts, client.TestOptions{})
}
