package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/hackernews/codegen/services"
)

func main() {
	err := services.Generate()
	if err != nil {
		log.Fatal(err)
	}
}
