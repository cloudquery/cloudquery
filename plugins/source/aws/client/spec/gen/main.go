package main

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/spec"
	"github.com/cloudquery/codegen/jsonschema"
)

func main() {
	fmt.Println("Generating JSON schema for plugin spec")
	jsonschema.GenerateIntoFile(new(spec.Spec), path.Join(currDir(), "..", "schema.json"),
		jsonschema.WithAddGoComments("github.com/cloudquery/cloudquery/plugins/source/aws/client/spec", path.Join(currDir(), "..")),
		jsonschema.WithAddGoComments("github.com/aws/aws-sdk-go-v2", path.Join(currDir(), "..", "..", "..", "vendor", "github.com/aws/aws-sdk-go-v2")),
	)
}

func currDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return path.Dir(filename)
}
