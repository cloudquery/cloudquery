package main

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/cloudquery/cloudquery/plugins/destination/test/v2/client"
	cqjsonschema "github.com/cloudquery/codegen/jsonschema"
	"github.com/invopop/jsonschema"
)

func main() {
	fmt.Println("Generating JSON schema for plugin spec")
	cqjsonschema.GenerateIntoFile(new(client.Spec), path.Join(currDir(), "../..", "schema.json"),
		cqjsonschema.WithAddGoComments("github.com/cloudquery/cloudquery/plugins/destination/test/v2/client", path.Join(currDir(), "../..")),
		func(reflector *jsonschema.Reflector) {
			reflector.NullableFromType = false
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
