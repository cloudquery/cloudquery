package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"

	"github.com/cockroachdb/cockroachdb-parser/pkg/sql/sem/tree"
	"github.com/mjibson/sqlfmt"
)

const policyResultsTable = `k8s_policy_results`

// Add example queries to the generated table docs from the policies/queries
// directory.
func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <path to queries directory> <path to table docs directory>", os.Args[0])
	}
	queriesDir := os.Args[1]
	tablesDir := os.Args[2]

	tables, err := readTables(tablesDir)
	if err != nil {
		log.Fatalf("error reading tables: %v", err)
	}
	queries, err := readQueries(queriesDir, tables)
	if err != nil {
		log.Fatalf("error reading queries: %v", err)
	}
	tableQueries := map[string][]query{}
	for _, q := range queries {
		if len(q.tables) == 0 {
			continue
		}
		for _, t := range q.tables {
			if _, ok := tableQueries[t]; !ok {
				tableQueries[t] = []query{}
			}
			tableQueries[t] = append(tableQueries[t], q)
		}
	}
	for t, qs := range tableQueries {
		var tbl table
		for _, tt := range tables {
			if tt.name == t {
				tbl = tt
				break
			}
		}
		err = addQueriesToTable(tbl, qs)
		if err != nil {
			log.Println("error adding queries to table:", err)
		}
	}
}

func addQueriesToTable(t table, qs []query) error {
	b, err := os.ReadFile(t.path)
	if err != nil {
		return err
	}
	s := string(b)
	if strings.Contains(s, "## Example Queries") {
		s = strings.Split(s, "## Example Queries")[0]
	}
	add := "\n\n## Example Queries\n\n"
	add += "These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.\n\n"
	entries := 0
	for _, q := range qs {
		if q.title == "" {
			// Skip queries without titles
			log.Println("Skipping query without title:", q.source)
			continue
		}
		smp, err := q.Simplify()
		if err != nil {
			log.Println("Skipping query due to error:", err)
			continue
		}
		add += fmt.Sprintf("### %s\n\n", q.title)
		add += fmt.Sprintf("```sql\n%s\n```\n\n", smp)
		entries++
	}
	add += "\n"
	if entries > 0 {
		s += add
	}
	return os.WriteFile(t.path, []byte(s), 0644)
}

type table struct {
	name string
	path string
}

func readTables(dir string) ([]table, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	ts := make([]table, 0, len(files))
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if filepath.Ext(f.Name()) != ".md" {
			continue
		}
		ts = append(ts, table{
			name: strings.TrimSuffix(f.Name(), ".md"),
			path: filepath.Join(dir, f.Name()),
		})
	}
	return ts, nil
}

type query struct {
	tables []string
	query  string
	title  string
	source string
}

func (q query) Simplify() (string, error) {
	qs := q.query
	qs = regexp.MustCompile(`[ \t]+`).ReplaceAllString(qs, " ")
	qs = regexp.MustCompile(`(?i)insert into `+policyResultsTable+`(?: \([^)]+\))?`).ReplaceAllString(qs, "")
	qs = strings.Trim(qs, "\n ")
	qs = regexp.MustCompile(`(?i):'execution_time'::timestamp as execution_time,`).ReplaceAllString(qs, "")
	qs = regexp.MustCompile(`(?i):'execution_time' as execution_time,`).ReplaceAllString(qs, "")
	qs = regexp.MustCompile(`(?i):'framework' as framework,`).ReplaceAllString(qs, "")
	qs = regexp.MustCompile(`(?i):'check_id' as check_id,`).ReplaceAllString(qs, "")
	n, err := sqlfmt.FmtSQL(tree.PrettyCfg{
		LineWidth:                80,
		TabWidth:                 2,
		DoNotNewLineAfterColName: false,
		Align:                    0,
		UseTabs:                  false,
		Simplify:                 true,
		Case:                     nil,
		JSONFmt:                  false,
	}, []string{qs})
	if err != nil {
		return "", err
	}
	return n, nil
}

func readQueries(dir string, tables []table) ([]query, error) {
	queries := []query{}
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		q, err := readQuery(path, tables)
		if err != nil {
			return err
		}
		queries = append(queries, q)
		return nil
	})
	return queries, err
}

func readQuery(path string, tables []table) (query, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return query{}, err
	}
	q := query{
		tables: []string{},
		query:  string(b),
		title:  extractTitleFromQuery(string(b)),
		source: path,
	}
	m := map[string]struct{}{}
	for _, t := range tables {
		if strings.Contains(q.query, t.name) {
			if _, ok := m[t.name]; ok {
				continue
			}
			m[t.name] = struct{}{}
			q.tables = append(q.tables, t.name)
		}
	}
	sort.Strings(q.tables)
	return q, nil
}

func extractTitleFromQuery(q string) string {
	r := regexp.MustCompile(`(?i)'(.+)'\s+as\s+title,`)
	matches := r.FindStringSubmatch(q)
	if len(matches) == 0 {
		fmt.Println("no matches for", q)
		return ""
	}
	title := matches[1]
	title = string(unicode.ToUpper(rune(title[0]))) + title[1:]
	return title
}
