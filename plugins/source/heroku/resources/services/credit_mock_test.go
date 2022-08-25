// Code generated by codegen; DO NOT EDIT.

package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	faker "github.com/cloudquery/faker/v3"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/julienschmidt/httprouter"
)

func createCredits() (client.HerokuService, error) {
	items := make(heroku.CreditListResult, 1)
	if err := faker.FakeData(&items); err != nil {
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
	s := heroku.NewService(ts.Client())
	s.URL = ts.URL
	return s, nil
}

func TestCredit(t *testing.T) {
	client.HerokuMockTestHelper(t, Credits(), createCredits, client.TestOptions{})
}
