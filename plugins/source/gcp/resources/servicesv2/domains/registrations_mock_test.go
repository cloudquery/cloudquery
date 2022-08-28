// Code generated by codegen; DO NOT EDIT.

package domains

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/domains/v1beta1"

	"google.golang.org/api/option"
)

type MockRegistrationsResult struct {
	Registrations []*domains.Registration `json:"registrations,omitempty"`
}

func createRegistrations() (*client.Services, error) {
	var item domains.Registration
	if err := faker.FakeData(&item); err != nil {
		return nil, err
	}

	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockRegistrationsResult{
			Registrations: []*domains.Registration{&item},
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
	svc, err := domains.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Domains: svc,
	}, nil
}

func TestRegistrations(t *testing.T) {
	client.MockTestHelper(t, Registrations(), createRegistrations, client.TestOptions{})
}
