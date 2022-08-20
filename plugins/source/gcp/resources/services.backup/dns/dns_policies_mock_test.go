package dns

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	dns "google.golang.org/api/dns/v1"
	"google.golang.org/api/option"
)

func createDnsPolicies() (*client.Services, error) {
	ctx := context.Background()
	var policy dns.Policy
	if err := faker.FakeData(&policy); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/dns/v1/projects/testProject/policies", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &dns.PoliciesListResponse{
			Policies: []*dns.Policy{&policy},
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
	svc, err := dns.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Dns: svc,
	}, nil
}

func TestDnsPolicies(t *testing.T) {
	client.GcpMockTestHelper(t, DNSPolicies(), createDnsPolicies, client.TestOptions{})
}
