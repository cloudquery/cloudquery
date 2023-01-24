package cohorts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/client"
	"github.com/cloudquery/cloudquery/plugins/source/mixpanel/internal/mixpanel"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createCohorts(router *mux.Router) error {
	var o mixpanel.Cohort
	if err := faker.FakeObject(&o); err != nil {
		return err
	}
	o.ProjectID = client.TestProjectID

	router.HandleFunc("/api/2.0/cohorts/list", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal([]mixpanel.Cohort{o})
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	var d struct {
		mixpanel.EngageProfileList
		Results []mixpanel.EngageProfileInResponse `json:"results"`
	}

	const numSamples = 7
	for i := 0; i < numSamples; i++ {
		var ep mixpanel.EngageProfileInResponse
		if err := faker.FakeObject(&ep); err != nil {
			return err
		}
		ep.DistinctID += "_" + strconv.FormatInt(int64(i), 10)
		ep.Properties = make(map[string]any)
		d.Results = append(d.Results, ep)
	}
	d.EngageProfileList.EngagePaginator.PageSize = 1000
	d.EngageProfileList.EngagePaginator.Total = int64(len(d.Results))
	d.EngageProfileList.EngagePaginator.SessionID = "the-session-id"
	d.CommonResponse.Status = mixpanel.StatusOK

	router.HandleFunc("/api/2.0/engage", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "unable to parse request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if fb := r.FormValue("filter_by_cohort"); fb != `{"id":123}` {
			http.Error(w, "invalid request: not a valid filter_by_cohort value: "+strconv.Quote(fb), http.StatusBadRequest)
			return
		}
		if pg := r.FormValue("page"); pg == "1" {
			fmt.Println("requested another page")
			if _, err := w.Write([]byte(`{"page":1}`)); err != nil {
				http.Error(w, "failed to write", http.StatusBadRequest)
			}
			return
		} else if pg != "" && pg != "0" {
			http.Error(w, "invalid request: not that many pages. requested page: "+strconv.Quote(pg), http.StatusBadRequest)
			return
		}

		b, err := json.Marshal(d)
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

func TestCohorts(t *testing.T) {
	client.MockTestHelper(t, Cohorts(), createCohorts, client.TestOptions{})
}
