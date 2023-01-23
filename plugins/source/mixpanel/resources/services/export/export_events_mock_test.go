package export

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createExportEvents(router *mux.Router) error {
	const numSamples = 77
	start := time.Date(2022, 2, 22, 22, 22, 22, 22, time.UTC)

	o := make([]mixpanel.ExportEvent, numSamples)
	for i := 0; i < numSamples; i++ {
		if err := faker.FakeObject(&o[i]); err != nil {
			return err
		}
		o[i].Properties = map[string]any{"time": start.Add(time.Duration(i) * time.Second).Unix(), "distinct_id": fmt.Sprintf("id%d", i)}
	}

	router.HandleFunc("/api/2.0/export", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < len(o); i++ {
			b, err := json.Marshal(o[i])
			if err != nil {
				http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
				return
			}
			if _, err := w.Write(b); err != nil {
				http.Error(w, "failed to write", http.StatusBadRequest)
				return
			}
			w.Write([]byte("\n"))
		}
	})

	return nil
}

func TestExportEvents(t *testing.T) {
	client.MockTestHelper(t, ExportEvents(), createExportEvents, client.TestOptions{})
}
