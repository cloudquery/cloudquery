//go:build mock
// +build mock

package crm

import (
	"context"
	"encoding/json"
	"github.com/cloudquery/cq-provider-gcp/client"
	faker "github.com/cloudquery/faker/v3"
	"github.com/julienschmidt/httprouter"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/option"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createCrmTestServer() (*client.Services, error) {
	ctx := context.Background()
	var project cloudresourcemanager.Project
	if err := faker.FakeData(&project); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/v3/projects/testProject", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(project)
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
	return &client.Services{
		Crm: svc,
	}, nil
}

func TestCrmProjects(t *testing.T) {
	client.GcpMockTestHelper(t, CrmProjects(), createCrmTestServer, client.TestOptions{})
}
