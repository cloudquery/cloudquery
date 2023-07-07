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

func createSubscriptionReservationRecommendations(router *mux.Router) error {
	var item armconsumption.LegacyReservationRecommendation
	if err := faker.FakeObject(&item); err != nil {
		return err
	}
	item.Kind = to.Ptr(armconsumption.ReservationRecommendationKindLegacy)

	resp := armconsumption.ReservationRecommendationsClientListResponse{
		ReservationRecommendationsListResult: armconsumption.ReservationRecommendationsListResult{
			Value: []armconsumption.ReservationRecommendationClassification{&item},
		},
	}
	resp.NextLink = to.Ptr("")
	router.HandleFunc("/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/reservationRecommendations", func(w http.ResponseWriter, r *http.Request) {
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

func TestSubscriptionReservationRecommendations(t *testing.T) {
	client.MockTestHelper(t, SubscriptionReservationRecommendations(), createSubscriptionReservationRecommendations)
}
