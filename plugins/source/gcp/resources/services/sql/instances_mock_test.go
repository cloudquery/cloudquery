package sql

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

func createInstances() (*client.Services, error) {
	var item sql.InstancesListResponse
	if err := faker.FakeObject(&item); err != nil {
		return nil, err
	}
	item.NextPageToken = ""

	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(item)
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
	svc, err := sql.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		SqlService: svc,
	}, nil
}

func TestInstances(t *testing.T) {
	client.MockTestHelper(t, Instances(), createInstances, client.TestOptions{})
}
