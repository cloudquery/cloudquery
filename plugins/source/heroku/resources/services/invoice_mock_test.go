package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/julienschmidt/httprouter"
)

func createInvoices() (*heroku.Service, error) {
	items := make(heroku.InvoiceListResult, 1)
	if err := faker.FakeObject(&items); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/*any", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		b, err := json.Marshal(items)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)
	c := *heroku.DefaultClient
	c.Transport = client.NewPaginator(ts.Client().Transport)
	s := heroku.NewService(&c)
	s.URL = ts.URL
	return s, nil
}

func TestInvoice(t *testing.T) {
	client.MockTestHelper(t, Invoices(), createInvoices, client.TestOptions{})
}
