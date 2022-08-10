package compute

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/faker/v3"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func createBackendServiceServer() (*client.Services, error) {
	ctx := context.Background()
	var backendService compute.BackendService
	if err := faker.FakeData(&backendService); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/projects/testProject/aggregated/backendServices", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &compute.BackendServiceAggregatedList{
			Items: map[string]compute.BackendServicesScopedList{
				"": {
					BackendServices: []*compute.BackendService{&backendService},
				},
			},
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
	svc, err := compute.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Compute: svc,
	}, nil
}

func TestComputeBackendServices(t *testing.T) {
	client.GcpMockTestHelper(t, ComputeBackendServices(), createBackendServiceServer, client.TestOptions{})
}
