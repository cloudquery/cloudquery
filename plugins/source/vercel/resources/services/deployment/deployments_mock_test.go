package deployment

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

func createDeployments(router *mux.Router) error {
	var vd vercel.Deployment
	if err := faker.FakeObject(&vd); err != nil {
		return err
	}
	t := vercel.MilliTime(time.Now())
	vd.CreatedAt = t
	vd.Ready = &t
	vd.BuildingAt = t

	var vc vercel.DeploymentCheck
	if err := faker.FakeObject(&vc); err != nil {
		return err
	}
	vc.CreatedAt = t
	vc.UpdatedAt = &t
	vc.StartedAt = &t
	vc.CompletedAt = &t

	router.HandleFunc("/v6/deployments", func(w http.ResponseWriter, r *http.Request) {
		list := struct {
			Deployments []vercel.Deployment `json:"deployments"`
			Pagination  vercel.Paginator    `json:"pagination"`
		}{
			Deployments: []vercel.Deployment{vd},
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

	router.HandleFunc("/v1/deployments/"+vd.UID+"/checks", func(w http.ResponseWriter, r *http.Request) {
		list := struct {
			Checks     []vercel.DeploymentCheck `json:"checks"`
			Pagination vercel.Paginator         `json:"pagination"`
		}{
			Checks: []vercel.DeploymentCheck{vc},
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

func TestDeployments(t *testing.T) {
	client.MockTestHelper(t, Deployments(), createDeployments)
}
