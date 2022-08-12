package monitoring

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	monitoring "google.golang.org/api/monitoring/v3"
	"google.golang.org/api/option"
)

func createMonitoringAlertPolicies() (*client.Services, error) {
	ctx := context.Background()
	var alertPolicy monitoring.AlertPolicy
	if err := faker.FakeData(&alertPolicy); err != nil {
		return nil, err
	}
	alertPolicy.Validity.Details = nil
	mux := httprouter.New()
	mux.GET("/v3/projects/testProject/alertPolicies", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &monitoring.ListAlertPoliciesResponse{
			AlertPolicies: []*monitoring.AlertPolicy{&alertPolicy},
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
	svc, err := monitoring.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Monitoring: svc,
	}, nil
}

func TestMonitoringAlertPolicies(t *testing.T) {
	client.GcpMockTestHelper(t, MonitoringAlertPolicies(), createMonitoringAlertPolicies, client.TestOptions{})
}
