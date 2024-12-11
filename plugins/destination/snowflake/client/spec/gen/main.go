package main

import (
	"fmt"
	"log"
	"path"
	"runtime"

	"github.com/cloudquery/cloudquery/plugins/destination/snowflake/v4/client"
	cqjsonschema "github.com/cloudquery/codegen/jsonschema"
)

func main() {
	fmt.Println("Generating JSON schema for plugin spec")
	cqjsonschema.GenerateIntoFile(new(client.Spec), path.Join(currDir(), "../..", "schema.json"),
		cqjsonschema.WithAddGoComments("github.com/cloudquery/cloudquery/plugins/destination/snowflake/v4/client", path.Join(currDir(), "../..")),
	)
}

func currDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return path.Dir(filename)
}
