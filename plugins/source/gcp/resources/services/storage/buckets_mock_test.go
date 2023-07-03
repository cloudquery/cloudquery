package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"cloud.google.com/go/iam"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	oldapi "google.golang.org/api/storage/v1"
)

func createBuckets(mux *httprouter.Router) error {
	var item oldapi.Bucket
	if err := faker.FakeObject(&item); err != nil {
		return fmt.Errorf("failed to fake data: %w", err)
	}
	item.TimeCreated = "2006-01-02T15:04:05Z07:00"
	item.Updated = "2006-01-02T15:04:05Z07:00"
	item.RetentionPolicy.EffectiveTime = "2006-01-02T15:04:05Z"

	var policy iam.Policy3
	if err := faker.FakeObject(&policy); err != nil {
		return err
	}

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

	return nil
}

func TestBuckets(t *testing.T) {
	client.MockTestRestHelper(t, Buckets(), createBuckets, client.TestOptions{})
}
