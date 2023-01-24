package engage

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

func createEngageRevenues(router *mux.Router) error {
	revData := make(map[string]mixpanel.EngageRevenue)
	start := time.Date(2022, 2, 22, 22, 22, 22, 22, time.UTC)

	const numSamples = 7
	for i := 0; i < numSamples; i++ {
		var rd mixpanel.EngageRevenue
		if err := faker.FakeObject(&rd); err != nil {
			return err
		}
		revData[start.Add(time.Duration(i)*86400*time.Second).Format("2006-01-02")] = rd
	}
	if len(revData) != numSamples {
		panic("failed to create revenue data")
	}
	revData["$overall"] = mixpanel.EngageRevenue{
		Amount:    99.99,
		Count:     10,
		PaidCount: 2,
	}

	router.HandleFunc("/api/2.0/engage/revenue", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(struct {
			Data any `json:"data"`
		}{
			Data: revData,
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

func TestEngageRevenues(t *testing.T) {
	client.MockTestHelper(t, EngageRevenues(), createEngageRevenues, client.TestOptions{})
}
