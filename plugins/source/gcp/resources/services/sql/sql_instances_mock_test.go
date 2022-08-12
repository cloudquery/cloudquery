package sql

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

func createSqlTestServer() (*client.Services, error) {
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
	return &client.Services{
		Sql: svc,
	}, nil
}

func TestSQLInstances(t *testing.T) {
	client.GcpMockTestHelper(t, SQLInstances(), createSqlTestServer, client.TestOptions{})
}
