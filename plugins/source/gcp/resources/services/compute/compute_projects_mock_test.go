package compute

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cloudquery/cq-provider-gcp/client"
	faker "github.com/cloudquery/faker/v3"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func createComputeProjects() (*client.Services, error) {
	id := "testProject"
	mux := httprouter.New()
	var project compute.Project
	if err := faker.FakeData(&project); err != nil {
		return nil, err
	}
	project.Name = id
	project.CreationTimestamp = time.Now().Format(time.RFC3339)

	mux.GET("/projects/testProject", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &project
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
	svc, err := compute.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Compute: svc,
	}, nil
}

func TestComputeProjects(t *testing.T) {
	client.GcpMockTestHelper(t, ComputeProjects(), createComputeProjects, client.TestOptions{})
}
