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
	domains "google.golang.org/api/domains/v1beta1"
	"google.golang.org/api/option"
)

func createDomainTestServer() (*client.Services, error) {
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
	return &client.Services{
		Domain: svc,
	}, nil
}

func TestDomainsRegistration(t *testing.T) {
	client.GcpMockTestHelper(t, DomainsRegistration(), createDomainTestServer, client.TestOptions{})
}
