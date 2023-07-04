package bigquery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/bigquery/v2"
)

func createBigqueryDatasets(mux *httprouter.Router) error {
	id := "testDataset"
	var dataset bigquery.Dataset
	if err := faker.FakeObject(&dataset); err != nil {
		return err
	}
	dataset.Id = id
	dataset.DatasetReference = &bigquery.DatasetReference{
		DatasetId: id,
	}
	mux.GET("/projects/testProject/datasets", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &bigquery.DatasetList{
			Datasets: []*bigquery.DatasetListDatasets{
				{
					DatasetReference: dataset.DatasetReference,
				},
			},
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

	mux.GET("/projects/testProject/datasets/testDataset", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(&dataset)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	mux.GET("/projects/testProject/datasets/testDataset/tables", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &bigquery.TableList{
			Tables: []*bigquery.TableListTables{
				{
					Id: id,
					TableReference: &bigquery.TableReference{
						TableId: id,
					},
				},
			},
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

	var table bigquery.Table
	if err := faker.FakeObject(&table); err != nil {
		return err
	}
	table.Id = id
	table.TableReference = &bigquery.TableReference{
		TableId: id,
	}
	schema := bigquery.TableSchema{
		Fields: []*bigquery.TableFieldSchema{{
			Name: "test",
			Type: "test",
		},
		},
	}
	table.Schema = &schema

	table.ExternalDataConfiguration = &bigquery.ExternalDataConfiguration{
		Autodetect: true,
		Schema:     &schema,
		SourceUris: []string{"test"},
	}
	table.Labels = map[string]string{
		"test": "test",
	}
	table.Clustering = &bigquery.Clustering{
		Fields: []string{"test"},
	}
	if err := faker.FakeObject(&table.Description); err != nil {
		return err
	}
	if err := faker.FakeObject(&table.EncryptionConfiguration); err != nil {
		return err
	}

	mux.GET("/projects/testProject/datasets/testDataset/tables/:table", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Println("what")
		fmt.Println(r.URL)
		b, err := json.Marshal(&table)
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

func TestBigqueryDatasets(t *testing.T) {
	client.MockTestRestHelper(t, Datasets(), createBigqueryDatasets, client.TestOptions{})
}
