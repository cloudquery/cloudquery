package auditlog

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/launchdarkly/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
	ldapi "github.com/launchdarkly/api-client-go/v11"
)

func createAuditLog(router *mux.Router) error {
	var o ldapi.AuditLogEntryListingRep
	if err := faker.FakeObject(&o); err != nil {
		return err
	}

	router.HandleFunc("/api/v2/auditlog", func(w http.ResponseWriter, r *http.Request) {
		list := ldapi.AuditLogEntryListingRepCollection{
			Items: []ldapi.AuditLogEntryListingRep{o},
		}

		b, err := json.Marshal(&list)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("content-type", "application/json")
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestAuditLogEntries(t *testing.T) {
	client.MockTestHelper(t, AuditLogEntries(), createAuditLog, client.TestOptions{})
}
