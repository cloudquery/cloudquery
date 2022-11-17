package main

import (
	"log"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/services"
)

func main() {
	err := services.Generate()
	if err != nil {
		log.Fatalln(err)
	}
}
