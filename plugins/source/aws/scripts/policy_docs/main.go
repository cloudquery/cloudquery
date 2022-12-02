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

var reInline = regexp.MustCompile(`(?i)\\ir (.+)`)
var reTable = regexp.MustCompile(`(?i)(?:from|join)\s+(\w+)`)
var reTitle = regexp.MustCompile(`(?i)\'(.+)\'\s+as\s+title,`)
var reIsPolicyQuery = regexp.MustCompile(`(?i)insert\s+into\s+aws_policy_results`)

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
	Path   string
	Tables []string
}

func getPolicyInfo(dir string, policy Policy) (*PolicyInfo, error) {
	queries, err := extractQueries(path.Join(dir, policy.Path))
	if err != nil {
		return nil, fmt.Errorf("failed to extract queries: %w", err)
	}
	return &PolicyInfo{
		Name:    policy.Name,
		Queries: queries,
	}, nil
}

func extractQueries(sqlPath string) ([]Query, error) {
	b, err := os.ReadFile(sqlPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %v: %w", sqlPath, err)
	}
	content := string(b)
	var queries []Query
	tableMatches := reTable.FindAllStringSubmatch(content, -1)
	isPolicyQuery := reIsPolicyQuery.MatchString(content)
	if isPolicyQuery && len(tableMatches) > 0 {
		q := Query{}
		for _, m := range tableMatches {
			q.Tables = append(q.Tables, m[1])
			q.Path = sqlPath
		}

		titleMatches := reTitle.FindAllStringSubmatch(content, -1)
		if len(titleMatches) == 0 {
			log.Printf("WARN: Failed to find title for query in %v", sqlPath)
		} else if len(titleMatches) >= 1 {
			q.Title = titleMatches[0][1]
		}
		queries = append(queries, q)
	}

	// recurse to find queries in inlined files
	dir := filepath.Dir(sqlPath)
	matches := reInline.FindAllStringSubmatch(content, -1)
	for _, m := range matches {
		q, err := extractQueries(path.Join(dir, m[1]))
		if err != nil {
			return nil, err
		}
		queries = append(queries, q...)
	}
	return queries, nil
}

func main() {
	if len(os.Args) <= 2 {
		log.Fatalf("Usage: %s <path to policies directory> <output dir>", os.Args[0])
	}
	dir := os.Args[1]
	index, err := readIndex(dir)
	if err != nil {
		log.Fatalf("error reading index: %v", err)
	}

	for _, p := range index.Policies {
		pi, err := getPolicyInfo(dir, p)
		if err != nil {
			log.Fatalf("error reading SQL for policy: %v", err)
		}
		fmt.Println(pi.Name, len(pi.Queries))
		for _, q := range pi.Queries {
			fmt.Println(q.Title, q.Tables, q.Path)
		}
	}
}
