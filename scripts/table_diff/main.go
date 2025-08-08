package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strings"

	"cloudquery/tablesdiff/changes"

	"github.com/bluekeyes/go-gitdiff/gitdiff"
)

var (
	sourcePluginDocsRegex = regexp.MustCompile(`^plugins/source/.*?/docs/tables/.*\.md$`)
)

func isPluginTableDocFile(file *gitdiff.File) bool {
	if file.IsBinary {
		return false
	}
	// Skip the README as we have everything we need from the tables files
	if strings.HasSuffix(file.OldName, "README.md") || strings.HasSuffix(file.NewName, "README.md") {
		return false
	}
	return sourcePluginDocsRegex.MatchString(file.OldName) || sourcePluginDocsRegex.MatchString(file.NewName)
}

func filterFiles(files []*gitdiff.File) []*gitdiff.File {
	filtered := make([]*gitdiff.File, 0)
	for _, file := range files {
		if isPluginTableDocFile(file) {
			filtered = append(filtered, file)
		}
	}
	return filtered
}

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("Usage: %s <diff-file-path> <changes-output-file>", os.Args[0])
	}
	diffFile := os.Args[1]
	outFile := os.Args[2]
	log.Printf("Reading diff file: %s", diffFile)
	patch, err := os.Open(diffFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Parsing diff file: %s", diffFile)
	files, _, err := gitdiff.Parse(patch)
	if err != nil {
		log.Fatal(err)
	}

	docsFiles := filterFiles(files)
	docChanges, err := changes.GetChanges(docsFiles)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Detected %d doc changes", len(docChanges))
	out, err := json.MarshalIndent(docChanges, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(outFile, out, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
