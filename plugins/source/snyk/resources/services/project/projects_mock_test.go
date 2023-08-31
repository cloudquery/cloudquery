package project

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func createProjects(mux *httprouter.Router) error {
	var project snyk.Project
	if err := faker.FakeObject(&project); err != nil {
		return err
	}
	mux.POST("/org/:orgID/projects", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		type projectsRoot struct {
			Projects []snyk.Project `json:"projects,omitempty"`
		}
		b, err := json.Marshal(projectsRoot{Projects: []snyk.Project{project}})
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

func TestProjects(t *testing.T) {
	client.MockTestHelper(t, Projects(), createProjects)
}
