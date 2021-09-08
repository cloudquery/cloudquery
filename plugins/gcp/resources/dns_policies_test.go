package resources

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/dns/v1"
	"google.golang.org/api/option"
)

func createDnsPolicies() (*dns.Service, error) {
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
	return svc, nil
}

func TestDnsPolicies(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: DNSPolicies(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			dnsSvc, err := createDnsPolicies()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Dns: dnsSvc,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, Provider, resource)
}
