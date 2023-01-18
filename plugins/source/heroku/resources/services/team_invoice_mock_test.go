package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/faker"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/julienschmidt/httprouter"
)

func createTeamInvoices() (*heroku.Service, error) {
	primaryItems := make(heroku.TeamListResult, 1)
	if err := faker.FakeObject(&primaryItems); err != nil {
		return nil, err
	}
	items := make(heroku.TeamInvoiceListResult, 1)
	if err := faker.FakeObject(&items); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	calls := 0
	mux.GET("/*any", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var b []byte
		var err error
		if calls == 0 {
			// return primary items on first call
			b, err = json.Marshal(primaryItems)
		} else {
			// return secondary items on subsequent calls
			b, err = json.Marshal(items)
		}
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
		calls++
	})
	ts := httptest.NewServer(mux)
	c := *heroku.DefaultClient
	c.Transport = client.NewPaginator(ts.Client().Transport)
	s := heroku.NewService(&c)
	s.URL = ts.URL
	return s, nil
}

func TestTeamInvoice(t *testing.T) {
	client.MockTestHelper(t, TeamInvoices(), createTeamInvoices, client.TestOptions{})
}
