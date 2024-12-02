package main

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/cloudquery/cloudquery/plugins/destination/kafka/v5/client/spec"
	cqjsonschema "github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/filetypes/v4"
	"github.com/invopop/jsonschema"
)

func main() {
	fmt.Println("Generating JSON schema for plugin spec")
	cqjsonschema.GenerateIntoFile(new(spec.Spec), path.Join(currDir(), "..", "schema.json"),
		append(filetypes.FileSpec{}.JSONSchemaOptions(),
			cqjsonschema.WithAddGoComments("github.com/cloudquery/cloudquery/plugins/destination/kafka/v5/client/spec", path.Join(currDir(), "..")),
			cqjsonschema.WithAddGoComments("github.com/cloudquery/filetypes/v4", path.Join(currDir(), "..", "..", "..", "vendor", "github.com/cloudquery/filetypes/v4")),
			func(r *jsonschema.Reflector) {
				// not required for this plugin
				r.NullableFromType = false
			},
		)...,
	)
}

func currDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return path.Dir(filename)
}
