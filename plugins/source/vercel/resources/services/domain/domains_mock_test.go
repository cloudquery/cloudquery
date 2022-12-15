package domain

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

func createDomains(router *mux.Router) error {
	var vd vercel.Domain
	if err := faker.FakeObject(&vd); err != nil {
		return err
	}
	t := vercel.MilliTime(time.Now())
	vd.CreatedAt = t
	vd.BoughtAt = &t
	vd.ConfigVerifiedAt = &t
	vd.ExpiresAt = &t
	vd.NsVerifiedAt = &t
	vd.OrderedAt = &t
	vd.TransferStartedAt = &t
	vd.TransferredAt = &t
	vd.TxtVerifiedAt = &t

	var vr vercel.DomainRecord
	if err := faker.FakeObject(&vr); err != nil {
		return err
	}
	vr.CreatedAt = &t
	vr.UpdatedAt = &t

	router.HandleFunc("/v5/domains", func(w http.ResponseWriter, r *http.Request) {
		list := struct {
			Domains    []vercel.Domain  `json:"domains"`
			Pagination vercel.Paginator `json:"pagination"`
		}{
			Domains: []vercel.Domain{vd},
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

	router.HandleFunc("/v4/domains/"+vd.Name+"/records", func(w http.ResponseWriter, r *http.Request) {
		list := struct {
			Records    []vercel.DomainRecord `json:"records"`
			Pagination vercel.Paginator      `json:"pagination"`
		}{
			Records: []vercel.DomainRecord{vr},
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

func TestDomains(t *testing.T) {
	client.MockTestHelper(t, Domains(), createDomains)
}
