package integration

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func createIntegrations(mux *httprouter.Router) error {
	const basePath = "/org/:orgID/integrations"
	mux.GET(basePath, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(snyk.Integrations{
			snyk.GitHubIntegrationType: uuid.NewString(),
		})
		if err != nil {
			http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var integration snyk.Integration
	if err := faker.FakeObject(&integration); err != nil {
		return err
	}

	mux.GET(basePath+"/"+snyk.GitHubIntegrationType, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(integration)
		if err != nil {
			http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var settings snyk.IntegrationSettings
	if err := faker.FakeObject(&settings); err != nil {
		return err
	}

	mux.GET(basePath+"/"+integration.ID+"/settings", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(settings)
		if err != nil {
			http.Error(w, "unable to marshal response: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestIntegrations(t *testing.T) {
	client.MockTestHelper(t, Integrations(), createIntegrations)
}
