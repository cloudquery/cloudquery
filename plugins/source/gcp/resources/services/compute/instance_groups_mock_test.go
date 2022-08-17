package compute

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
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func createInstanceGroups() (*client.Services, error) {
	ctx := context.Background()
	var instanceGroup compute.InstanceGroup
	if err := faker.FakeData(&instanceGroup); err != nil {
		return nil, err
	}
	instanceGroup.Zone = "test/test"
	instanceGroup.CreationTimestamp = "2021-06-02T14:13:54.470Z"
	mux := httprouter.New()
	mux.GET("/projects/testProject/aggregated/InstanceGroups", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &compute.InstanceGroupAggregatedList{
			Items: map[string]compute.InstanceGroupsScopedList{
				"": {
					InstanceGroups: []*compute.InstanceGroup{&instanceGroup},
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

	var inst compute.InstanceWithNamedPorts
	if err := faker.FakeData(&inst); err != nil {
		return nil, err
	}
	mux.POST(fmt.Sprintf("/projects/testProject/zones/test/instanceGroups/%s/listInstances", instanceGroup.Name), func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &compute.InstanceGroupsListInstances{
			Items: []*compute.InstanceWithNamedPorts{&inst},
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

func TestComputeInstanceGroups(t *testing.T) {
	client.GcpMockTestHelper(t, InstanceGroups(), createInstanceGroups, client.TestOptions{})
}
