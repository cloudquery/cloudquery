package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cloudquery/cq-provider-k8s/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/migration"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func main() {
	fmt.Println("BEGIN;")
	for r, t := range provider.Provider().ResourceMap {
		fmt.Printf("-- %s\n", r)
		up, err := migration.CreateTableDefinitions(context.Background(), schema.PostgresDialect{}, t, nil)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error creating tables for %s: %v\n", r, err)
			os.Exit(1)
		}
		fmt.Println(strings.Join(up, "\n"))
	}
	fmt.Println("COMMIT;")
}
