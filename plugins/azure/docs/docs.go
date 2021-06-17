package main

import (
	"fmt"

	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	err := docs.GenerateDocs(resources.Provider(), "./docs")
	if err != nil {
		fmt.Sprintf("Failed to geneerate docs: %s", err)
	}
}