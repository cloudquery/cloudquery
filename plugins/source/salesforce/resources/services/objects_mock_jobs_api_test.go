package services

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/salesforce/client"
	"github.com/gorilla/mux"
)

func createObjectsBulkAPI(router *mux.Router) error {
	describeRes := describeResponse{
		Fields: []map[string]any{
			{
				"name": "Id",
				"type": "string",
			},
			{
				"name": "Name",
				"type": "string",
			},
		},
	}

	router.HandleFunc("/services/data/v56.0/sobjects/{object}/describe", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&describeRes)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var createQueryJobRes createQueryJobResponse
	createQueryJobRes.Id = "7501x00001"

	router.HandleFunc("/services/data/v56.0/jobs/query", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&createQueryJobRes)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var getQueryJobStatusRes getQueryJobStatusResponse
	getQueryJobStatusRes.State = StateJobComplete
	getQueryJobStatusRes.Id = "7501x00001"

	router.HandleFunc("/services/data/v56.0/jobs/query/{queryId}", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&getQueryJobStatusRes)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	jobResponse := `Id,Name
0011x00001,Test
0011x00002,Test2
`
	router.HandleFunc("/services/data/v56.0/jobs/query/{queryId}/results", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte(jobResponse)); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	return nil
}

func TestObjectsBulkAPI(t *testing.T) {
	client.MockTestHelper(t, Objects(), createObjectsBulkAPI)
}
