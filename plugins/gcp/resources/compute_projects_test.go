package resources

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func TestComputeProjects(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: ComputeProjects(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			computeSvc, err := createComputeProjects()
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
	providertest.TestResource(t, Provider, resource)
}

func createComputeProjects() (*compute.Service, error) {
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
	return svc, nil
}
