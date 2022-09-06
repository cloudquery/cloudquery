package main

import (
	"fmt"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/plugin"
	"github.com/cloudquery/plugin-sdk/docs"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	outputPath := "./docs"
	dir, err := ioutil.ReadDir(outputPath + "/tables")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
		os.Exit(1)
	}
	for _, d := range dir {
		if err := os.RemoveAll(path.Join([]string{outputPath + "/tables", d.Name()}...)); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
			os.Exit(1)
		}

	}
	if err = docs.GenerateSourcePluginDocs(plugin.Plugin(), path.Join(outputPath, "tables")); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
	}
}
