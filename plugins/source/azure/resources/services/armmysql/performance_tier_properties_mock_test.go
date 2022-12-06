// Code generated by codegen; DO NOT EDIT.

package armmysql

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
)

func createPerformanceTierProperties(router *mux.Router) error {
	var item armmysql.LocationBasedPerformanceTierClientListResponse
	if err := faker.FakeObject(&item); err != nil {
		return err
	}

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/performanceTiers", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&item)
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

func TestPerformanceTierProperties(t *testing.T) {
	client.MockTestHelper(t, PerformanceTierProperties(), createPerformanceTierProperties)
}
