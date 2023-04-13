package reporting

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v2/faker"
	"github.com/julienschmidt/httprouter"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func createIssues(mux *httprouter.Router) error {
	var resp snyk.ListReportingIssuesResponse
	if err := faker.FakeObject(&resp); err != nil {
		return err
	}
	resp.Total = 2001
	i := 0
	mux.POST("/v1/reporting/issues/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp.Results[0].Issue.ID = fmt.Sprintf("test-%d", i)
		i++
		b, err := json.Marshal(resp)
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

func TestIssues(t *testing.T) {
	client.MockTestHelper(t, Issues(), createIssues)
}
