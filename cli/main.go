package main

import (
	"os"

	"github.com/cloudquery/cloudquery/cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
