package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/resources/plugin"
	"log"

	"github.com/cloudquery/plugin-sdk/docs"
)

func main() {
	err := docs.GenerateSourcePluginDocs(plugin.Plugin(), "./docs")
	if err != nil {
		log.Fatalf("Failed to generate docs: %s", err)
	}
}
