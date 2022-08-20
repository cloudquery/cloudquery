package compute

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func createComputeURLMaps() (*client.Services, error) {
	ctx := context.Background()
	var maps compute.UrlMapList
	if err := faker.FakeData(&maps); err != nil {
		return nil, err
	}
	maps.NextPageToken = ""
	mux := httprouter.New()
	mux.GET("/projects/testProject/global/urlMaps", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(maps)
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
	svc, err := compute.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Compute: svc,
	}, nil
}

func TestComputeURLMaps(t *testing.T) {
	client.GcpMockTestHelper(t, ComputeURLMaps(), createComputeURLMaps, client.TestOptions{})
}
