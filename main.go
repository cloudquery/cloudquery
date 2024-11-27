// CloudQuery uses a monorepo approach with a separate Go module per component. Visit our GitHub repository to see the full project structure and all the components https://github.com/cloudquery/cloudquery
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
