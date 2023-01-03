package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/codegen/recipes"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/stripe/stripe-go/v74"
)

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	testdataDir := path.Join(path.Dir(filename), "..", "..", "resources", "testdata")

	fixture := struct {
		Resources map[string]any `json:"resources"`
	}{
		Resources: make(map[string]any),
	}

	extras := []any{
		&stripe.BankAccount{},
		&stripe.SubscriptionItem{},
		&stripe.LineItem{},
		&stripe.TaxID{},
	}

	renames := map[string]string{
		"invoice_item": "invoiceitem",
	}

	for _, e := range extras {
		recipes.AllResources = append(recipes.AllResources, &recipes.Resource{
			DataStruct: e,
			Service:    "extras", // only to make GenerateNames happy
		})
	}

	for _, r := range recipes.AllResources {
		r.Infer()
		r.GenerateNames()

		ds := r.DataStruct
		if err := faker.FakeObject(ds, faker.WithMaxDepth(6)); err != nil {
			log.Fatal(err)
		}

		// Some structs reference other structs which exceeds max depth at one point, and the whole field is left with default value
		switch item := ds.(type) {
		case *stripe.Customer:
			item.DefaultSource = &stripe.PaymentSource{ID: "test"}
		case *stripe.Invoice:
			item.DefaultSource = &stripe.PaymentSource{ID: "test"}
		case *stripe.Subscription:
			item.DefaultSource = &stripe.PaymentSource{ID: "test"}
		}

		keyName := r.SubService
		if n := renames[keyName]; n != "" {
			keyName = n
		}

		fixture.Resources[keyName] = ds
	}

	b, _ := json.MarshalIndent(fixture, "", "  ")

	fn := path.Join(testdataDir, "fixtures_gen.json")
	log.Println("Writing", fn)
	if err := os.WriteFile(fn, b, 0644); err != nil {
		log.Fatal(err)
	}
}
