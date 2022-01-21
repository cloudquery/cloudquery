package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cq-provider-gcp/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/migration"
)

func main() {
	if err := migration.Run(context.Background(), provider.Provider(), ""); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(1)
	}
}
