package main

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/cloudquery/cloudquery/plugins/destination/azblob/v4/client/spec"
	"github.com/cloudquery/codegen/jsonschema"
	"github.com/cloudquery/filetypes/v4"
)

func main() {
	fmt.Println("Generating JSON schema for plugin spec")
	jsonschema.GenerateIntoFile(new(spec.Spec), path.Join(currDir(), "..", "schema.json"),
		append(filetypes.FileSpec{}.JSONSchemaOptions(),
			jsonschema.WithAddGoComments("github.com/cloudquery/cloudquery/plugins/destination/azblob/v4/client/spec", path.Join(currDir(), "..")),
			jsonschema.WithAddGoComments("github.com/cloudquery/filetypes/v4", path.Join(currDir(), "..", "..", "..", "vendor", "github.com/cloudquery/filetypes/v4")),
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
