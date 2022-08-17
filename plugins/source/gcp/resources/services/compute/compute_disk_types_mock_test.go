package compute

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func createDiskTypeServer() (*client.Services, error) {
	ctx := context.Background()
	var diskType compute.DiskType
	if err := faker.FakeData(&diskType); err != nil {
		return nil, err
	}
	diskType.SelfLink = "https://www.googleapis.com/compute/v1/projects/project-id/zones/zone-id/disks/disk-id"
	mux := httprouter.New()
	mux.GET("/projects/testProject/aggregated/diskTypes", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &compute.DiskTypeAggregatedList{
			Items: map[string]compute.DiskTypesScopedList{
				"": {
					DiskTypes: []*compute.DiskType{&diskType},
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

func TestComputeDiskTypes(t *testing.T) {
	client.GcpMockTestHelper(t, ComputeDiskTypes(), createDiskTypeServer, client.TestOptions{})
}
