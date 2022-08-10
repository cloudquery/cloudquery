package cloudfunctions

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/faker/v3"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/cloudfunctions/v1"
	"google.golang.org/api/option"
)

func createCloudFunctionsTestServer() (*client.Services, error) {
	ctx := context.Background()
	var function cloudfunctions.CloudFunction
	if err := faker.FakeData(&function); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/*data", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &cloudfunctions.ListFunctionsResponse{
			Functions: []*cloudfunctions.CloudFunction{&function},
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
	svc, err := cloudfunctions.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		CloudFunctions: svc,
	}, nil
}

func TestCloudfunctionsFunction(t *testing.T) {
	client.GcpMockTestHelper(t, CloudfunctionsFunction(), createCloudFunctionsTestServer, client.TestOptions{})
}
