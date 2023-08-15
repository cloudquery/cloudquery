package reporting

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func createIssuesForPath(path string) func(mux *httprouter.Router) error {
	return func(mux *httprouter.Router) error {
		resp := snyk.ListReportingIssuesResponse{Total: 2001}
		for i := 0; i < resp.Total; i++ {
			var result snyk.ListReportingIssueResult
			if err := faker.FakeObject(&result); err != nil {
				return err
			}
			result.Issue.ID = fmt.Sprintf("test-%d", i)
			resp.Results = append(resp.Results, result)
		}

		mux.POST(path, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			page, err := strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				http.Error(w, "unable to parse page", http.StatusBadRequest)
			}
			perPage, err := strconv.Atoi(r.URL.Query().Get("perPage"))
			if err != nil {
				http.Error(w, "unable to parse perPage", http.StatusBadRequest)
			}
			start := (page - 1) * perPage
			end := page * perPage
			if end > resp.Total {
				end = resp.Total
			}
			pageResults := resp.Results[start:end]
			b, err := json.Marshal(snyk.ListReportingIssuesResponse{Total: 2001, Results: pageResults})
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
}

func TestIssues(t *testing.T) {
	client.MockTestHelper(t, Issues(), createIssuesForPath("/v1/reporting/issues/"))
}
