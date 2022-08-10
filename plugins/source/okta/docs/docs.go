package main

import (
	"log"

	"github.com/cloudquery/cq-provider-okta/resources"
	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	err := docs.GenerateDocs(resources.Provider(), "./docs", true)
	if err != nil {
		log.Fatalf("Failed to geneerate docs: %s", err)
	}
}
