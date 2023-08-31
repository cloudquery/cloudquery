package security

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/azure/client"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createSettings(router *mux.Router) error {
	var d armsecurity.DataExportSettings
	if err := faker.FakeObject(&d); err != nil {
		return err
	}

	item := armsecurity.SettingsClientListResponse{
		SettingsList: armsecurity.SettingsList{
			Value: []armsecurity.SettingClassification{
				&d,
			},
		},
	}

	emptyStr := ""
	item.NextLink = &emptyStr

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Security/settings", func(w http.ResponseWriter, r *http.Request) {
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

func TestSettings(t *testing.T) {
	client.MockTestHelper(t, Settings(), createSettings)
}
