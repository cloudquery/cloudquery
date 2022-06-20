package main

import (
	"log"
	"os"

	"github.com/cloudquery/cloudquery/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
