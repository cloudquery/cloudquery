package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"cloud.google.com/go/iam"
	"cloud.google.com/go/storage"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	oldapi "google.golang.org/api/storage/v1"
)

func createBuckets() (*client.Services, error) {
	var item oldapi.Bucket
	if err := faker.FakeObject(&item); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	item.TimeCreated = "2006-01-02T15:04:05Z07:00"
	item.Updated = "2006-01-02T15:04:05Z07:00"
	item.RetentionPolicy.EffectiveTime = "2006-01-02T15:04:05Z"

	var policy iam.Policy3
	if err := faker.FakeObject(&policy); err != nil {
		return nil, err
	}

	mux := httprouter.New()
	mux.GET("/b", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&oldapi.Buckets{
			Items: []*oldapi.Bucket{&item},
		})
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	mux.GET("/b/test string/iam", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	svc, err := storage.NewClient(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		StorageClient: svc,
	}, nil
}

func TestBuckets(t *testing.T) {
	client.MockTestHelper(t, Buckets(), createBuckets, client.TestOptions{})
}
