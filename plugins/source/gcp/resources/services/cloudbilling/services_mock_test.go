package cloudbilling

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/cloudbilling/v1"
	"google.golang.org/api/option"
)

func createServicesTestServer() (*client.Services, error) {
	ctx := context.Background()
	var service cloudbilling.Service
	if err := faker.FakeData(&service); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/v1/services", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &cloudbilling.ListServicesResponse{
			Services: []*cloudbilling.Service{&service},
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

	var sku cloudbilling.Sku
	if err := faker.FakeData(&sku); err != nil {
		return nil, err
	}
	mux.GET(fmt.Sprintf("/v1/%s/skus", service.Name), func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &cloudbilling.ListSkusResponse{
			Skus: []*cloudbilling.Sku{&sku},
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
	svc, err := cloudbilling.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		CloudBilling: svc,
	}, nil
}

func TestServices(t *testing.T) {
	client.GcpMockTestHelper(t, Services(), createServicesTestServer, client.TestOptions{})
}
