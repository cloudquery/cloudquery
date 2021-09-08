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
	"google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/option"
)

func createResourceManagerProjects() (*cloudresourcemanager.Service, error) {
	ctx := context.Background()
	var project cloudresourcemanager.Project
	if err := faker.FakeData(&project); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	project.CreateTime = time.Now().Format(time.RFC3339)
	project.DeleteTime = time.Now().Format(time.RFC3339)
	project.UpdateTime = time.Now().Format(time.RFC3339)
	project.ProjectId = "testProject"
	mux.GET("/v3/projects/testProject", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&project)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var policy cloudresourcemanager.Policy
	if err := faker.FakeData(&policy); err != nil {
		return nil, err
	}
	mux.POST("/v3/projects/testProject:getIamPolicy", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(policy)
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
	svc, err := cloudresourcemanager.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func TestResourceManagerProjects(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: ResourceManagerProjects(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			resourceManager, err := createResourceManagerProjects()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				ResourceManager: resourceManager,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, Provider, resource)
}
