package main

import (
	"fmt"

	"github.com/cloudquery/cq-provider-azure/resources/provider"

	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	err := docs.GenerateDocs(provider.Provider(), "./docs", true)
	if err != nil {
		fmt.Sprintf("Failed to geneerate docs: %s", err)
	}
}
