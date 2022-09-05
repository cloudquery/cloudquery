package main

import (
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/plugins/source/terraform/resources/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	err := docs.GenerateDocs(provider.Provider(), "./docs", true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
	}
}
