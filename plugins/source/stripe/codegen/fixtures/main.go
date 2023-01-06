package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/stripe/codegen/recipes"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/stripe/stripe-go/v74"
)

var dataStructs map[string]any // pkgPath vs. struct

func main() {
	extras := []any{
		&stripe.BankAccount{},
		&stripe.CreditNoteLineItem{},
		&stripe.FeeRefund{},
		&stripe.LineItem{},
		&stripe.SubscriptionItem{},
		&stripe.TaxID{},
		&stripe.TransferReversal{},
	}

	renames := map[string]string{
		"invoice_item": "invoiceitem",
	}

	prefixReplaces := map[string]string{
		"billing_portal_": "billing_portal.",
		"checkout_":       "checkout.",
		"identity_":       "identity.",
		"issuing_":        "issuing.",
		"radar_":          "radar.",
		"reporting_":      "reporting.",
		"sigma_":          "",
		"terminal_":       "terminal.",
		"treasury_":       "treasury.",
	}

	copies := map[string]string{
		"line_item": "item",
	}

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

	dataStructs = make(map[string]any)

	for _, r := range recipes.AllResources {
		traverseStructs(r.DataStruct, "", 0)
	}
	for _, extra := range extras {
		traverseStructs(extra, "", 0)
	}

	csr := caser.New()
	for _, ds := range dataStructs {
		ds := ds
		if err := faker.FakeObject(&ds, faker.WithMaxDepth(6)); err != nil {
			log.Fatal(err)
		}

		typ := reflect.TypeOf(ds)
		for typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}

		keyName := csr.ToSnake(typ.Name())
		if n := renames[keyName]; n != "" {
			keyName = n
		} else {
			for k, v := range prefixReplaces {
				if strings.HasPrefix(keyName, k) {
					keyName = v + keyName[len(k):]
					break
				}
			}
		}

		if c, ok := copies[keyName]; ok {
			fixture.Resources[c] = ds
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

func traverseStructs(ds any, allowedPath string, depth int) {
	if depth > 16 {
		return
	}
	typ := reflect.TypeOf(ds)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return
	}

	if allowedPath == "" {
		allowedPath = typ.PkgPath()
	} else {
		if typ.PkgPath() != allowedPath {
			return
		}
	}

	mapKey := typ.PkgPath() + "." + typ.Name()

	if _, ok := dataStructs[mapKey]; ok {
		return
	}

	dataStructs[mapKey] = ds

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		traverseStructs(reflect.New(f.Type).Interface(), allowedPath, depth+1)
	}
}
