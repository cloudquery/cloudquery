package iam

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	iam "google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
)

func createIamRolesTestServer() (*client.Services, error) {
	ctx := context.Background()
	var acc iam.ServiceAccount
	if err := faker.FakeData(&acc); err != nil {
		return nil, err
	}
	acc.Name = "test"
	mux := httprouter.New()

	var role iam.Role
	if err := faker.FakeData(&role); err != nil {
		return nil, err
	}
	mux.GET("/v1/projects/testProject/roles", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &iam.ListRolesResponse{
			Roles: []*iam.Role{&role},
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
	svc, err := iam.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Iam: svc,
	}, nil
}

func TestIamRoles(t *testing.T) {
	client.GcpMockTestHelper(t, IamRoles(), createIamRolesTestServer, client.TestOptions{})
}
