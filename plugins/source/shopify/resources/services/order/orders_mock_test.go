package order

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createOrders(router *mux.Router) error {
	var so shopify.Order
	if err := faker.FakeObject(&so); err != nil {
		return err
	}

	router.HandleFunc("/admin/api/"+shopify.APIVersion+"/orders.json", func(w http.ResponseWriter, r *http.Request) {
		list := shopify.GetOrdersResponse{
			Orders:   []shopify.Order{so},
			PageSize: 1,
		}

		b, err := json.Marshal(&list)
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

func TestOrders(t *testing.T) {
	client.MockTestHelper(t, Orders(), createOrders, client.TestOptions{})
}
