package serviceusage

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
	"google.golang.org/api/serviceusage/v1"
)

func createServices() (*client.Services, error) {
	ctx := context.Background()
	var service serviceusage.GoogleApiServiceusageV1Service
	if err := faker.FakeDataSkipFields(&service, []string{"Config"}); err != nil {
		return nil, err
	}
	service.Config = &serviceusage.GoogleApiServiceusageV1ServiceConfig{}
	if err := faker.FakeDataSkipFields(service.Config, []string{"Documentation"}); err != nil {
		return nil, err
	}
	service.Config.Documentation = &serviceusage.Documentation{}
	service.Config.Apis[0].Methods[0].Options[0].Value = []byte("{}")
	service.Config.Apis[0].Options[0].Value = []byte("{}")
	if err := faker.FakeDataSkipFields(service.Config.Documentation, []string{"Pages"}); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/v1/projects/testProject/services", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &serviceusage.ListServicesResponse{
			Services: []*serviceusage.GoogleApiServiceusageV1Service{&service},
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
	svc, err := serviceusage.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		ServiceUsage: svc,
	}, nil
}

func TestServices(t *testing.T) {
	client.GcpMockTestHelper(t, Services(), createServices, client.TestOptions{})
}
