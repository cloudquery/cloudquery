package resource_manager

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v3"
	"google.golang.org/api/option"
)

func createResourceManagerFolders() (*client.Services, error) {
	ctx := context.Background()
	var folder cloudresourcemanager.Folder
	if err := faker.FakeData(&folder); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	folder.CreateTime = time.Now().Format(time.RFC3339)
	folder.DeleteTime = time.Now().Format(time.RFC3339)
	folder.UpdateTime = time.Now().Format(time.RFC3339)
	folder.Name = "testProject"
	mux.GET("/v3/folders", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &cloudresourcemanager.ListFoldersResponse{
			Folders: []*cloudresourcemanager.Folder{&folder},
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

	var policy cloudresourcemanager.Policy
	if err := faker.FakeData(&policy); err != nil {
		return nil, err
	}
	mux.POST("/v3/folders/testProject:getIamPolicy", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	return &client.Services{
		ResourceManager: svc,
	}, nil
}

func TestResourceManagerFolders(t *testing.T) {
	client.GcpMockTestHelper(t, ResourceManagerFolders(), createResourceManagerFolders, client.TestOptions{})
}
