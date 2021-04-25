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
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
)

func TestIamRoles(t *testing.T) {
	resource := ResourceTestData{
		Table: resources.IamRoles(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
			Resources:  []client.Resource{{Name: "iam.project_roles"}},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			iamSvc, err := createIamTestServer()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Iam: iamSvc,
			})
			return c, nil
		},
	}
	testResource(t, resources.Provider, resource)
}
func TestIamServiceAccounts(t *testing.T) {
	resource := ResourceTestData{
		Table: resources.IamServiceAccounts(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
			Resources:  []client.Resource{{Name: "iam.service_accounts"}},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			iamSvc, err := createIamTestServer()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Iam: iamSvc,
			})
			return c, nil
		},
	}
	testResource(t, resources.Provider, resource)
}

func createIamTestServer() (*iam.Service, error) {
	ctx := context.Background()
	var acc iam.ServiceAccount
	if err := faker.FakeData(&acc); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/v1/projects/testProject/serviceAccounts", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &iam.ListServiceAccountsResponse{
			Accounts: []*iam.ServiceAccount{&acc},
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
	return svc, nil
}
