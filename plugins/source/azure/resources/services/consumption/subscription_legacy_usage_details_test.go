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

func createSubscriptionLegacyUsageDetails(router *mux.Router) error {
	var (
		item1 armconsumption.LegacyUsageDetail
		item2 armconsumption.ModernUsageDetail
	)
	if err := faker.FakeObject(&item1); err != nil {
		return err
	}
	item1.Kind = to.Ptr(armconsumption.UsageDetailsKindLegacy)

	if err := faker.FakeObject(&item2); err != nil {
		return err
	}
	item2.Kind = to.Ptr(armconsumption.UsageDetailsKindModern) // will be skipped, but API seems to return both

	resp := armconsumption.UsageDetailsClientListResponse{
		UsageDetailsListResult: armconsumption.UsageDetailsListResult{
			// Value is an interface{} so we can't mock it directly
			Value: []armconsumption.UsageDetailClassification{&item1, &item2},
		},
	}
	resp.NextLink = to.Ptr("")

	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{periodName}/providers/Microsoft.Consumption/usageDetails", func(w http.ResponseWriter, r *http.Request) {
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

func TestSubscriptionLegacyUsageDetails(t *testing.T) {
	client.MockTestHelper(t, SubscriptionLegacyUsageDetails(), createSubscriptionLegacyUsageDetails)
}
