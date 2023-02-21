package team

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

func createTeams(router *mux.Router) error {
	var vt vercel.Team
	if err := faker.FakeObject(&vt); err != nil {
		return err
	}
	t := vercel.MilliTime(time.Now())
	vt.CreatedAt = t
	vt.UpdatedAt = &t

	var vm vercel.TeamMember
	if err := faker.FakeObject(&vm); err != nil {
		return err
	}
	vm.CreatedAt = t
	vm.AccessRequestedAt = &t

	router.HandleFunc("/v2/teams", func(w http.ResponseWriter, r *http.Request) {
		list := struct {
			Teams      []vercel.Team    `json:"teams"`
			Pagination vercel.Paginator `json:"pagination"`
		}{
			Teams: []vercel.Team{vt},
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

	router.HandleFunc("/v2/teams/"+vt.ID+"/members", func(w http.ResponseWriter, r *http.Request) {
		list := struct {
			Members    []vercel.TeamMember `json:"members"`
			Pagination vercel.Paginator    `json:"pagination"`
		}{
			Members: []vercel.TeamMember{vm},
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

func TestTeams(t *testing.T) {
	client.MockTestHelper(t, Teams(), createTeams, client.TestOptions{})
}
