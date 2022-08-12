package memorystore

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	"google.golang.org/api/redis/v1"
)

func createInstancesServer() (*client.Services, error) {
	ctx := context.Background()
	var inst redis.Instance
	if err := faker.FakeData(&inst); err != nil {
		return nil, err
	}
	inst.Name = "projects/{project_id}/locations/{location_id}/instances/{instance_id}"
	mux := httprouter.New()
	mux.GET("/*data", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &redis.ListInstancesResponse{
			Instances: []*redis.Instance{&inst},
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
	svc, err := redis.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Redis: svc,
	}, nil
}

func TestRedisInstances(t *testing.T) {
	client.GcpMockTestHelper(t, RedisInstances(), createInstancesServer, client.TestOptions{})
}
