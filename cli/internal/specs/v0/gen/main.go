package main

import (
	"fmt"
	"log"
	"path"
	"reflect"
	"runtime"

	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	cqgen_jsonschema "github.com/cloudquery/codegen/jsonschema"
	"github.com/invopop/jsonschema"
)

func main() {
	fmt.Println("Generating JSON schema for CLI spec")
	specsType := reflect.TypeOf(specs.Spec{})
	cqgen_jsonschema.GenerateIntoFile(new(specs.Spec), path.Join(currDir(), "..", "schema.json"),
		cqgen_jsonschema.WithAddGoComments("github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0", path.Join(currDir(), "..")),
		func(r *jsonschema.Reflector) {
			r.AdditionalFields = func(t reflect.Type) []reflect.StructField {
				if t == specsType { // we need to add the extra fields, as the `spec` field is just `any`
					return reflect.VisibleFields(reflect.TypeOf(struct {
						Source      specs.Source
						Destination specs.Destination
						Transformer specs.Transformer
					}{}))
				}
				return nil
			}
		},
	)
}

func currDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return path.Dir(filename)
}
