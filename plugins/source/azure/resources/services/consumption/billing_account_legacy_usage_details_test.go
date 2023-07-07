package consumption

import (
	"encoding/json"
	"net/http"

	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createBillingAccountLegacyUsageDetails(router *mux.Router) error {
	var item armconsumption.LegacyUsageDetail
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	item.Kind = to.Ptr(armconsumption.UsageDetailsKindLegacy)

	resp := armconsumption.UsageDetailsClientListResponse{
		UsageDetailsListResult: armconsumption.UsageDetailsListResult{
			// Value is an interface{} so we can't mock it directly
			Value: []armconsumption.UsageDetailClassification{&item},
		},
	}
	resp.NextLink = to.Ptr("")

	router.HandleFunc("/providers/Microsoft.Billing/billingAccounts/"+client.LegacyAccountName+"/providers/Microsoft.Consumption/usageDetails", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(&resp)
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

func TestBillingAccountLegacyUsageDetails(t *testing.T) {
	client.MockTestHelper(t, BillingAccountLegacyUsageDetails(), createBillingAccountLegacyUsageDetails)
}
