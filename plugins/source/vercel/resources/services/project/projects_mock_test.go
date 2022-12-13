package project

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createProjects(router *mux.Router) error {
	var vp vercel.Project
	if err := faker.FakeObject(&vp); err != nil {
		return err
	}
	t := vercel.MilliTime(time.Now())
	vp.CreatedAt = t
	vp.UpdatedAt = t
	vp.TransferStartedAt = &t
	vp.TransferCompletedAt = &t

	var ve vercel.ProjectEnv
	if err := faker.FakeObject(&ve); err != nil {
		return err
	}
	ve.CreatedAt = t
	ve.UpdatedAt = &t

	router.HandleFunc("/v9/projects", func(w http.ResponseWriter, r *http.Request) {
		list := struct {
			Projects   []vercel.Project `json:"projects"`
			Pagination vercel.Paginator `json:"pagination"`
		}{
			Projects: []vercel.Project{vp},
		}

		b, err := json.Marshal(&list)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	router.HandleFunc("/v9/projects/"+vp.ID+"/env", func(w http.ResponseWriter, r *http.Request) {
		list := struct {
			Envs       []vercel.ProjectEnv `json:"envs"`
			Pagination vercel.Paginator    `json:"pagination"`
		}{
			Envs: []vercel.ProjectEnv{ve},
		}

		b, err := json.Marshal(&list)
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

func TestProjects(t *testing.T) {
	client.MockTestHelper(t, Projects(), createProjects)
}
