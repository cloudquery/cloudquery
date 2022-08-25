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

func createLogDrains() (client.HerokuService, error) {
	primaryItems := make(heroku.AppListResult, 1)
	if err := faker.FakeData(&primaryItems); err != nil {
		return nil, err
	}
	items := make(heroku.LogDrainListResult, 1)
	if err := faker.FakeData(&items); err != nil {
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
	s := heroku.NewService(ts.Client())
	s.URL = ts.URL
	return s, nil
}

func TestLogDrain(t *testing.T) {
	client.HerokuMockTestHelper(t, LogDrains(), createLogDrains, client.TestOptions{})
}
