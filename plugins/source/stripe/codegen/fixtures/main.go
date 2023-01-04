package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"

	"github.com/cloudquery/plugin-sdk/caser"
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

	dataStructs := []any{
		&stripe.Account{},
		&stripe.Customer{},
		&stripe.Dispute{},
		&stripe.Invoice{},
		&stripe.InvoiceItem{},
		&stripe.Product{},
		&stripe.Refund{},
		&stripe.Subscription{},
		&stripe.BankAccount{},
		&stripe.SubscriptionItem{},
		&stripe.LineItem{},
		&stripe.TaxID{},
	}

	renames := map[string]string{
		"invoice_item": "invoiceitem",
	}

	csr := caser.New()
	for _, ds := range dataStructs {
		ds := ds
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

		typ := reflect.TypeOf(ds)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}
		keyName := csr.ToSnake(typ.Name())
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
