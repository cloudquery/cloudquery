package logging

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	logging "google.golang.org/api/logging/v2"
	"google.golang.org/api/option"
)

func createLoggingMetrics() (*client.Services, error) {
	ctx := context.Background()
	var logMetric logging.LogMetric
	if err := faker.FakeData(&logMetric); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/v2/projects/testProject/metrics", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &logging.ListLogMetricsResponse{
			Metrics: []*logging.LogMetric{&logMetric},
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
	svc, err := logging.NewService(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Logging: svc,
	}, nil
}

func TestLoggingMetrics(t *testing.T) {
	client.GcpMockTestHelper(t, LoggingMetrics(), createLoggingMetrics, client.TestOptions{})
}
