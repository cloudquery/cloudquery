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
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func TestComputeInstances(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: resources.ComputeInstances(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			computeSvc, err := createInstancesServer()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Compute: computeSvc,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
}

func createInstancesServer() (*compute.Service, error) {
	ctx := context.Background()
	var inst compute.Instance
	if err := faker.FakeData(&inst); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/projects/testProject/aggregated/instances", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &compute.InstanceAggregatedList{
			Items: map[string]compute.InstancesScopedList{
				"": {
					Instances: []*compute.Instance{&inst},
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
	return svc, nil
}
