package customer

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createCustomers(router *mux.Router) error {
	var sc shopify.Customer
	if err := faker.FakeObject(&sc); err != nil {
		return err
	}

	router.HandleFunc("/admin/api/"+shopify.APIVersion+"/customers.json", func(w http.ResponseWriter, r *http.Request) {
		list := shopify.GetCustomersResponse{
			Customers: []shopify.Customer{sc},
			PageSize:  1,
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

func TestCustomers(t *testing.T) {
	client.MockTestHelper(t, Customers(), createCustomers, client.TestOptions{})
}
