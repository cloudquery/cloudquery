package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/exp/maps"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

var reInline = regexp.MustCompile(`(?im)^\\ir (.+)`)
var reTable = regexp.MustCompile(`(?i)(?:from|join)\s+(\w+)`)
var reTitle = regexp.MustCompile(`(?i)\:'check_id'(?:\s+as\s+check_id\s*)?,\s+'(.+)'(?:\s+as\s+title\s*)?,`)
var reViewName = regexp.MustCompile(`(?i)create\s+(or\s+replace\s+)view\s+(.+)\s+as`)
var reIsPolicyQuery *regexp.Regexp

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
	Tables  []string

	CreatedViews   map[string]bool
	DependentViews []string
	UnusedViews    []string
}

func newPolicyInfo(name string, queries []Query, allTables []Table) *PolicyInfo {
	pi := &PolicyInfo{
		Name:    name,
		Queries: removeDuplicates(queries),
	}
	pi.setTables(allTables)
	return pi
}

// Tables returns the unique set of tables needed to run all queries
func (pi *PolicyInfo) setTables(allTables Tables) {
	t, v := map[string]struct{}{}, map[string]struct{}{}
	pi.CreatedViews = map[string]bool{}

	// Add all views detected from CREATE VIEW statements
	for _, q := range pi.Queries {
		if q.View {
			pi.CreatedViews[q.ViewName] = true
		}
	}

	// Accessed views from queries, should be a subset of the above
	for _, q := range pi.Queries {
		for _, vi := range q.Views {
			v[vi] = struct{}{}
		}
	}

	for _, q := range pi.Queries {
		// Add all tables from queries *and* views
		for _, table := range q.Tables {
			if pi.CreatedViews[table] {
				// This is a view, not a table
				v[table] = struct{}{}
			}

			t[table] = struct{}{}
			ancestors := allTables.FindAncestors(table)
			for _, a := range ancestors {
				t[a.Name] = struct{}{}
			}
		}
	}

	pi.Tables, pi.DependentViews = maps.Keys(t), maps.Keys(v)

	unused := make(map[string]struct{}, len(pi.CreatedViews))
	for k := range pi.CreatedViews {
		if _, ok := v[k]; !ok {
			unused[k] = struct{}{}
		}
	}
	pi.UnusedViews = maps.Keys(unused)

	sort.Strings(pi.Tables)
	sort.Strings(pi.DependentViews)
	sort.Strings(pi.UnusedViews)
}

type Query struct {
	Title  string // Empty for views
	Path   string
	Tables []string
	Views  []string

	View     bool   `json:",omitempty"`
	ViewName string `json:",omitempty"`
}

func removeDuplicates(queries []Query) []Query {
	var clean []Query
	m := map[string]bool{}
	for _, q := range queries {
		if _, found := m[q.Path]; found {
			continue
		}
		m[q.Path] = true
		clean = append(clean, q)
	}
	return clean
}

func getPolicyInfo(prefix string, tables Tables, dir string, policy Policy) (*PolicyInfo, error) {
	queries, err := extractQueries(prefix, path.Join(dir, policy.Path))
	if err != nil {
		return nil, fmt.Errorf("failed to extract queries: %w", err)
	}
	return newPolicyInfo(policy.Name, removeDuplicates(queries), tables), nil
}

func extractQueries(prefix, sqlPath string) ([]Query, error) {
	b, err := os.ReadFile(sqlPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %v: %w", sqlPath, err)
	}
	content := string(b)
	var queries []Query
	q := Query{}
	if tableMatches := reTable.FindAllStringSubmatch(content, -1); len(tableMatches) > 0 {
		for _, m := range tableMatches {
			switch {
			case strings.HasPrefix(m[1], prefix):
				// This could be a table or still a view, will sort out later
				q.Tables = append(q.Tables, m[1])
				q.Path = sqlPath
			case strings.HasPrefix(m[1], "view_"+prefix):
				q.Views = append(q.Views, m[1])
				q.Path = sqlPath
			}
		}
	}

	isPolicyQuery := reIsPolicyQuery.MatchString(content)
	if isPolicyQuery && (len(q.Tables)+len(q.Views)) > 0 {
		titleMatches := reTitle.FindAllStringSubmatch(content, -1)
		if len(titleMatches) == 0 {
			return nil, fmt.Errorf("failed to find title for query in %v", sqlPath)
		} else if len(titleMatches) >= 1 {
			q.Title = titleMatches[0][1]
		}
		queries = append(queries, q)
	} else {
		mv := reViewName.FindStringSubmatch(content)
		if len(mv) > 0 {
			q.View = true
			q.ViewName = mv[2]
			if !strings.HasPrefix(q.ViewName, prefix+"_") && !strings.HasPrefix(q.ViewName, "view_"+prefix) {
				return nil, fmt.Errorf("view %q in %s does not start with `%s_` or `view_%s`", q.ViewName, sqlPath, prefix, prefix)
			}
			queries = append(queries, q)
		}
	}

	// recurse to find queries in inlined files
	dir := filepath.Dir(sqlPath)
	matches := reInline.FindAllStringSubmatch(content, -1)
	for _, m := range matches {
		q, err := extractQueries(prefix, path.Join(dir, m[1]))
		if err != nil {
			return nil, err
		}
		queries = append(queries, q...)
	}
	return queries, nil
}

func writePolicyDocs(info []*PolicyInfo, outputPath string) error {
	t, err := template.New("policies.md.go.tpl").
		Funcs(template.FuncMap{
			"add": func(a int, b int) int {
				return a + b
			},
		}).
		ParseFS(templatesFS, "templates/policies.md.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse template for README.md: %v", err)
	}
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file %v: %v", outputPath, err)
	}
	defer f.Close()
	if err := t.Execute(f, info); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}
	return nil
}

type Table struct {
	Name      string `json:"name"`
	Relations Tables `json:"relations"`
}

type Tables []Table

// FindAncestors returns the list of ancestors for a given table, if any.
func (tt Tables) FindAncestors(name string) []Table {
	for _, t := range tt {
		if t.Name == name {
			return []Table{}
		}
		r := t.Relations.FindAncestors(name)
		if r != nil {
			return append([]Table{t}, r...)
		}
	}
	return nil
}

func readTablesJSON(filepath string) ([]Table, error) {
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read json file: %w", err)
	}
	var tables []Table
	err = json.Unmarshal(b, &tables)
	return tables, err
}

func main() {
	if len(os.Args) <= 4 {
		log.Fatalf("Usage: %s <table prefix> <path to policies directory> <output filename> <path to __tables.json>", os.Args[0])
	}
	prefix := os.Args[1]
	dir := os.Args[2]
	out := os.Args[3]
	tablesPath := os.Args[4]
	index, err := readIndex(dir)
	if err != nil {
		log.Fatalf("error reading index: %v", err)
	}
	tables, err := readTablesJSON(tablesPath)
	if err != nil {
		log.Fatalf("error reading tables JSON: %v", err)
	}

	reIsPolicyQuery = regexp.MustCompile(`(?i)insert\s+into\s+` + prefix + `_policy_results`)

	var info []*PolicyInfo
	for _, p := range index.Policies {
		pi, err := getPolicyInfo(prefix, tables, dir, p)
		if err != nil {
			log.Fatalf("error reading policy info: %v", err)
		}
		info = append(info, pi)
	}
	err = writePolicyDocs(info, out)
	if err != nil {
		log.Fatalf("failed to write policy documentation to %v: %v", out, err)
	}
}
