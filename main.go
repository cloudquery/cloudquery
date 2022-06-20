package main

import (
	"os"

	"github.com/cloudquery/cloudquery/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
