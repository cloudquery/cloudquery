package iam

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	iam "google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
)

func createIamServiceAccountsTestServer() (*client.Services, error) {
	ctx := context.Background()
	var acc iam.ServiceAccount
	if err := faker.FakeData(&acc); err != nil {
		return nil, err
	}
	acc.Name = "test"
	mux := httprouter.New()
	mux.GET("/v1/projects/testProject/serviceAccounts", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &iam.ListServiceAccountsResponse{
			Accounts: []*iam.ServiceAccount{&acc},
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
	if err := faker.FakeData(&key); err != nil {
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
	svc, err := iam.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Iam: svc,
	}, nil
}

func TestIamServiceAccounts(t *testing.T) {
	client.GcpMockTestHelper(t, IamServiceAccounts(), createIamServiceAccountsTestServer, client.TestOptions{})
}
