package main

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/client/spec"
	"github.com/cloudquery/codegen/jsonschema"
)

func main() {
	fmt.Println("Generating JSON schema for plugin spec")
	jsonschema.GenerateIntoFile(new(spec.Spec), path.Join(currDir(), "..", "schema.json"),
		jsonschema.WithAddGoComments("github.com/cloudquery/cloudquery/plugins/destination/clickhouse/v6/client/spec", path.Join(currDir(), "..")),
	)
}

func currDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return path.Dir(filename)
}
