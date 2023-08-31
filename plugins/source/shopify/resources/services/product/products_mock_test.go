package product

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/client"
	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/shopify"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/gorilla/mux"
)

func createProducts(router *mux.Router) error {
	var sp shopify.Product
	if err := faker.FakeObject(&sp); err != nil {
		return err
	}
	sp.Tags = shopify.Tags{"tag1", "tag2"}

	router.HandleFunc("/admin/api/"+shopify.APIVersion+"/products.json", func(w http.ResponseWriter, r *http.Request) {
		list := shopify.GetProductsResponse{
			Products: []shopify.Product{sp},
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

func TestProducts(t *testing.T) {
	client.MockTestHelper(t, Products(), createProducts, client.TestOptions{})
}
