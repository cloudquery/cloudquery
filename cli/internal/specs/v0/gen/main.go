package main

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	"github.com/cloudquery/codegen/jsonschema"
)

func main() {
	fmt.Println("Generating JSON schema for CLI spec")
	jsonschema.GenerateIntoFile(new(specs.Spec), path.Join(currDir(), "..", "schema.json"),
		jsonschema.WithAddGoComments("github.com/cloudquery/cloudquery/cli/internal/specs/v0", path.Join(currDir(), "..")),
	)
}

func currDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return path.Dir(filename)
}
