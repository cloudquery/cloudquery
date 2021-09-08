package resources

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cq-provider-gcp/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/monitoring/v3"
	"google.golang.org/api/option"
)

func createMonitoringAlertPolicies() (*monitoring.Service, error) {
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
	return svc, nil
}

func TestMonitoringAlertPolicies(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: MonitoringAlertPolicies(),
		Config: client.Config{
			ProjectIDs: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			monitoringSvc, err := createMonitoringAlertPolicies()
			if err != nil {
				return nil, err
			}
			c := client.NewGcpClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"}, &client.Services{
				Monitoring: monitoringSvc,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, Provider, resource)
}
