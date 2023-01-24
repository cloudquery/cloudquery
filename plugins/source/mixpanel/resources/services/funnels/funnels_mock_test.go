package funnels

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createFunnels(router *mux.Router) error {
	var o mixpanel.Funnel
	if err := faker.FakeObject(&o); err != nil {
		return err
	}

	router.HandleFunc("/api/2.0/funnels/list", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal([]mixpanel.Funnel{o})
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	funnelData := make(map[string]mixpanel.FunnelData)
	start := time.Date(2022, 2, 22, 22, 22, 22, 22, time.UTC)

	const numSamples = 7
	for i := 0; i < numSamples; i++ {
		var fd mixpanel.FunnelData
		if err := faker.FakeObject(&fd); err != nil {
			return err
		}
		fd.Analysis = map[string]any{"test": i}
		funnelData[start.Add(time.Duration(i)*86400*time.Second).Format("2006-01-02")] = fd
	}
	if len(funnelData) != numSamples {
		panic("failed to create funnel data")
	}

	router.HandleFunc("/api/2.0/funnels", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(struct {
			Data any `json:"data"`
		}{
			Data: funnelData,
		})
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

func TestFunnels(t *testing.T) {
	client.MockTestHelper(t, Funnels(), createFunnels, client.TestOptions{})
}
