package main

import (
	"github.com/cloudquery/cq-provider-gcp/resources/provider"
	"log"

	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	err := docs.GenerateDocs(provider.Provider(), "./docs", true)
	if err != nil {
		log.Fatalf("Failed to geneerate docs: %s", err)
	}
}
