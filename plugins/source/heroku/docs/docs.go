package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	outputPath := "./docs"
	dir, err := os.ReadDir(outputPath + "/tables")
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
	// TODO: migrate to SDK v2
	//if err = docs.GenerateSourcePluginDocs(plugin.Plugin(), path.Join(outputPath, "tables")); err != nil {
	//	fmt.Fprintf(os.Stderr, "Failed to generate docs: %s\n", err)
	//}
}
