package services

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/salesforce/client"
	"github.com/gorilla/mux"
)

func createObjectsQueryAPI(router *mux.Router) error {
	var queryRes queryResponse
	queryRes.TotalSize = 1
	queryRes.Done = true
	queryRes.Records = make([]map[string]any, 1)
	queryRes.Records[0] = make(map[string]any)
	queryRes.Records[0]["Id"] = "0011x00001"
	queryRes.Records[0]["Name"] = "Test"

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
			{
				"name": "test",
				"type": "base64",
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

	router.HandleFunc("/services/data/v56.0/query", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&queryRes)
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

func TestObjectsQueryAPI(t *testing.T) {
	client.MockTestHelper(t, Objects(), createObjectsQueryAPI)
}
