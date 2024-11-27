// This top level main file is intended so we show up in https://pkg.go.dev/github.com/cloudquery/cloudquery
// We publish the CLI from the cli sub package
package main

import (
	"context"
	"log"

	"github.com/cloudquery/cloudquery/cli/cmd"
)

func main() {
	err := cmd.NewCmdRoot().ExecuteContext(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
