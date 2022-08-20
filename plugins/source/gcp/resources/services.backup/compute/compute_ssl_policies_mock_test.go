package compute

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func createSslPolicies() (*client.Services, error) {
	ctx := context.Background()
	var p compute.SslPolicy
	if err := faker.FakeData(&p); err != nil {
		return nil, err
	}
	p.CreationTimestamp = time.Now().Format(time.RFC3339)
	mux := httprouter.New()
	mux.GET("/projects/testProject/global/SslPolicies", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &compute.SslPoliciesList{
			Items: []*compute.SslPolicy{
				&p,
			},
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
	svc, err := compute.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Compute: svc,
	}, nil
}

func TestComputeSslPolicies(t *testing.T) {
	client.GcpMockTestHelper(t, ComputeSslPolicies(), createSslPolicies, client.TestOptions{})
}
