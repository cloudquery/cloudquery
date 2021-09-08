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
	"google.golang.org/api/option"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

func createSqlTestServer() (*sql.Service, error) {
	ctx := context.Background()
	var db sql.DatabaseInstance
	if err := faker.FakeData(&db); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/sql/v1beta4/projects/:project/instances", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &sql.InstancesListResponse{
			Items: []*sql.DatabaseInstance{&db},
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
	svc, err := sql.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return svc, nil
}

func TestSQLInstances(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: resources.SQLInstances(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			sqlSvc, err := createSqlTestServer()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Sql: sqlSvc,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
}
