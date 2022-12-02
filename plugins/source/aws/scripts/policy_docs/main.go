package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

var reInline = regexp.MustCompile(`\\ir (.+)`)
var reTable = regexp.MustCompile(`(?:from|join)\s+(\w+)`)

type Index struct {
	Policies []Policy `json:"policies"`
}

type Policy struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func readIndex(dir string) (*Index, error) {
	indexFile := path.Join(dir, "index.json")
	content, err := os.ReadFile(indexFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", indexFile, err)
	}
	var index Index
	err = json.Unmarshal(content, &index)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json from index file %s: %w", indexFile, err)
	}
	return &index, nil
}

type PolicyInfo struct {
	Name    string
	Queries []Query
}

type Query struct {
	Title  string
	Tables []string
}

func getPolicyInfo(policy Policy) (*PolicyInfo, error) {
	queries, err := extractQueries(policy.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to extract queries: %w", err)
	}
	return &PolicyInfo{
		Name:    policy.Name,
		Queries: queries,
	}, nil
}

func extractQueries(sqlPath string) ([]Query, error) {
	dir := filepath.Dir(sqlPath)
	content, err := os.ReadFile(sqlPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %v: %w", sqlPath, err)
	}
	var queries []Query

}

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("Usage: %s <path to policies directory> <output dir>", os.Args[0])
	}
	dir := os.Args[1]
	index, err := readIndex(dir)
	if err != nil {
		log.Fatalf("error reading index: %v", err)
	}

	for _, p := range index.Policies {
		policyInfo := getPolicyInfo(p)
		fmt.Println(policyInfo)
	}

}
