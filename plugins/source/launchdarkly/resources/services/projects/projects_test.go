package projects

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
	ldapi "github.com/launchdarkly/api-client-go/v11"
)

func createProjects(router *mux.Router) error {
	var o ldapi.Project
	if err := faker.FakeObject(&o); err != nil {
		return err
	}

	router.HandleFunc("/api/v2/projects", func(w http.ResponseWriter, r *http.Request) {
		list := ldapi.Projects{
			Items: []ldapi.Project{o},
		}

		b, err := json.Marshal(&list)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var m ldapi.MetricListingRep
	if err := faker.FakeObject(&m); err != nil {
		return err
	}

	router.HandleFunc("/api/v2/metrics/"+o.Name, func(w http.ResponseWriter, r *http.Request) {
		list := ldapi.MetricCollectionRep{
			Items: []ldapi.MetricListingRep{m},
		}

		b, err := json.Marshal(&list)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var fd ldapi.FlagDefaultsRep
	if err := faker.FakeObject(&fd); err != nil {
		return err
	}

	router.HandleFunc("/api/v2/projects/"+o.Name+"/flag-defaults", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&fd)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var ff ldapi.FeatureFlag
	if err := faker.FakeObject(&ff); err != nil {
		return err
	}

	router.HandleFunc("/api/v2/flags/"+o.Name, func(w http.ResponseWriter, r *http.Request) {
		list := ldapi.FeatureFlags{
			Items: []ldapi.FeatureFlag{ff},
		}

		b, err := json.Marshal(&list)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestProjects(t *testing.T) {
	client.MockTestHelper(t, Projects(), createProjects, client.TestOptions{})
}
