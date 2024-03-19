package main

import (
	"fmt"
	"github.com/cloudquery/cloudquery/plugins/source/hackernews/client"
	"log"
	"path"
	"runtime"

	cqjsonschema "github.com/cloudquery/codegen/jsonschema"
)

func main() {
	fmt.Println("Generating JSON schema for plugin spec")
	cqjsonschema.GenerateIntoFile(new(client.Spec), path.Join(currDir(), "../..", "schema.json"),
		cqjsonschema.WithAddGoComments("github.com/cloudquery/cloudquery/plugins/source/hackernews/client", path.Join(currDir(), "../..")),
	)
}

func currDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return path.Dir(filename)
}
