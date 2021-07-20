package resources_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-gcp/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	domains "google.golang.org/api/domains/v1beta1"
	"google.golang.org/api/option"
)

func TestDomainsRegistration(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: resources.DomainsRegistration(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			domainSvc, err := createDomainTestServer()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Domain: domainSvc,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
}

func createDomainTestServer() (*domains.Service, error) {
	ctx := context.Background()
	var reg domains.Registration
	if err := faker.FakeData(&reg); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/v1beta1/projects/testProject/locations/-/registrations", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &domains.ListRegistrationsResponse{
			Registrations: []*domains.Registration{&reg},
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
	svc, err := domains.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return svc, nil
}
