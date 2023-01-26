package services

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/plausible/client"
	"github.com/gorilla/mux"
)

func createTimeSeriesService(router *mux.Router) error {
	response := StatsTimeseriesResponse{
		Results: []StatsTimeseriesResult{
			{
				Date:          "2021-01-01",
				Visitors:      1,
				PageViews:     1,
				VisitDuration: 1,
				BounceRate:    1,
				Visits:        1,
			},
		},
	}
	router.HandleFunc("/api/v1/stats/timeseries", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&response)
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

func TestStatsTimeseries(t *testing.T) {
	client.MockTestHelper(t, StatsTimeseries(), createTimeSeriesService)
}
