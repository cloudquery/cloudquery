package main

import (
	"fmt"
	"os"

	"github.com/cloudquery/cq-provider-sdk/provider/docs"
	"github.com/cloudquery/cq-provider-terraform/resources"
)

func main() {
	err := docs.GenerateDocs(resources.Provider(), "./docs")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
	}
}
