// CloudQuery uses a monorepo approach with a separate Go module per component. Visit our GitHub repository to see the full project structure and all the components https://github.com/cloudquery/cloudquery
package main

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/cli/v6/cmd"
)

func main() {
	if err := cmd.NewCmdRoot().ExecuteContext(context.Background()); err != nil {
		fmt.Println(err)
	}
}
