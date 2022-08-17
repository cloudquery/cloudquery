package storage

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
	storage "google.golang.org/api/storage/v1"
)

func createStorageTestServer() (*client.Services, error) {
	ctx := context.Background()
	var bucket storage.Bucket
	if err := faker.FakeData(&bucket); err != nil {
		return nil, err
	}
	bucket.Name = "testBucket"
	mux := httprouter.New()
	mux.GET("/b", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &storage.Buckets{Items: []*storage.Bucket{&bucket}}
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

	var policy storage.Policy
	if err := faker.FakeData(&policy); err != nil {
		return nil, err
	}
	mux.GET("/b/testBucket/iam", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &policy
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
	svc, err := storage.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Storage: svc,
	}, nil
}

func TestStorageBucket(t *testing.T) {
	client.GcpMockTestHelper(t, StorageBuckets(), createStorageTestServer, client.TestOptions{})
}
