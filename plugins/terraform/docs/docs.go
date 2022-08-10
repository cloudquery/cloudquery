package main

import (
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/plugins/terraform/resources"
	"github.com/cloudquery/cq-provider-sdk/provider/docs"
)

func main() {
	err := docs.GenerateDocs(resources.Provider(), "./docs", true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
	}
}
