package serviceusage

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/serviceusage/v1"

	"google.golang.org/api/option"
)

type MockServicesResult struct {
	Services []*serviceusage.GoogleApiServiceusageV1Service `json:"services,omitempty"`
}

func createServices() (*client.Services, error) {
	var item serviceusage.GoogleApiServiceusageV1Service
	if err := faker.FakeDataSkipFields(&item, []string{"Config"}); err != nil {
		return nil, err
	}
	item.Config = &serviceusage.GoogleApiServiceusageV1ServiceConfig{}
	if err := faker.FakeDataSkipFields(item.Config, []string{"Documentation"}); err != nil {
		return nil, err
	}
	item.Config.Documentation = &serviceusage.Documentation{}
	item.Config.Apis[0].Methods[0].Options[0].Value = []byte("{}")
	item.Config.Apis[0].Options[0].Value = []byte("{}")
	if err := faker.FakeDataSkipFields(item.Config.Documentation, []string{"Pages"}); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockServicesResult{
			Services: []*serviceusage.GoogleApiServiceusageV1Service{&item},
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
	svc, err := serviceusage.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Serviceusage: svc,
	}, nil
}

func TestServices(t *testing.T) {
	client.MockTestHelper(t, Services(), createServices, client.TestOptions{})
}
