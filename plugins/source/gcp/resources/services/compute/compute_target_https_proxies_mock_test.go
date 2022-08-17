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

func createTargetHTTPSProxies() (*client.Services, error) {
	ctx := context.Background()
	var proxy compute.TargetHttpsProxy
	if err := faker.FakeData(&proxy); err != nil {
		return nil, err
	}
	proxy.CreationTimestamp = time.Now().Format(time.RFC3339)
	mux := httprouter.New()
	mux.GET("/projects/testProject/global/targetHttpsProxies", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &compute.TargetHttpsProxyList{
			Items: []*compute.TargetHttpsProxy{
				&proxy,
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

func TestComputeTargetHTTPSProxies(t *testing.T) {
	client.GcpMockTestHelper(t, ComputeTargetHTTPSProxies(), createTargetHTTPSProxies, client.TestOptions{})
}
