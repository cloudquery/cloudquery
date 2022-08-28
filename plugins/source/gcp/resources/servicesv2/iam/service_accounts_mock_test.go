// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/iam/v1"

	"google.golang.org/api/option"
)

type MockServiceAccountsResult struct {
	Accounts []*iam.ServiceAccount `json:"accounts,omitempty"`
}

func createServiceAccounts() (*client.Services, error) {
	var item iam.ServiceAccount
	if err := faker.FakeData(&item); err != nil {
		return nil, err
	}

	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockServiceAccountsResult{
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
