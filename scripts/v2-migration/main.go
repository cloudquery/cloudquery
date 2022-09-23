package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/cloudquery/cloudquery/scripts/v2-migration/internal/convert"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/gertd/go-pluralize"
)

var (
	outputFile   string
	v1TablesPath string
	v2TablesPath string
)

//go:embed templates/output.go.tpl
var outputTemplate string

var CustomTableComments = map[string]string{}

var RenamedTables = map[string]string{
	"aws_accounts": "aws_iam_accounts",
}

type Table struct {
	Name    string
	Columns []Column
	Status  string
	Comment string
}

type Column struct {
	Name    string
	Type    string
	Status  string
	Comment string
}

func readTables(dir string) []Table {
	items, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	tables := make([]Table, 0, 100)
	for _, it := range items {
		if strings.HasSuffix(it.Name(), ".md") && it.Name() != "README.md" {
			tables = append(tables, Table{
				Name:    strings.TrimSuffix(it.Name(), ".md"),
				Columns: readColumns(path.Join(dir, it.Name())),
			})
		}
	}
	return tables
}

func readColumns(file string) []Column {
	b, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	content := string(b)
	lines := strings.Split(content, "\n")
	inTable := false
	cols := make([]Column, 0)
	for _, line := range lines {
		if strings.HasPrefix(line, "| ------------- |") {
			inTable = true
			continue
		}
		if inTable && strings.HasPrefix(line, "|") {
			values := strings.Split(line, "|")
			name := values[1]
			name = strings.TrimSuffix(name, " (PK)")
			cols = append(cols, Column{
				Name: name,
				Type: values[2],
			})
		}
	}
	return cols
}

func compareTables(v1, v2 []Table) []Table {
	comparison := make(map[string]Table, 0)
	v1Map := tablesToMap(v1)
	v2Map := tablesToMap(v2)
	for name, t1 := range v1Map {
		t2, found := v2Map[name]
		if found {
			t1.Status = "updated"
			t1.Columns = compareColumns(t1, t2)
			comparison[name] = t1
		} else {
			t1.Status = "removed"
			replacement, foundReplacement := findLikelyTableReplacement(t1, v2)
			if foundReplacement {
				t1.Status = replacement.Status
				t1.Comment = replacement.Comment
			}
			if t1.Status == "renamed" {
				other := v2Map[replacement.Name]
				other.Status = "renamed"
				other.Comment = fmt.Sprintf("Renamed from [%s](%s)", t1.Name, t1.Name)
				other.Columns = compareColumns(t1, other)
				v2Map[replacement.Name] = other
			}

			comparison[name] = t1
		}
	}

	for name, t2 := range v2Map {
		_, found := v1Map[name]
		switch {
		case found:
			continue
		case t2.Status == "":
			t2.Status = "added"
			for c := range t2.Columns {
				t2.Columns[c].Status = "added"
			}
			comparison[name] = t2
		default:
			comparison[name] = t2
		}
	}

	tables := make([]Table, 0, len(comparison))
	for _, t := range comparison {
		if comment, ok := CustomTableComments[t.Name]; ok {
			t.Comment = comment
		}
		tables = append(tables, t)
	}
	sort.Slice(tables, func(i, j int) bool {
		return tables[i].Name < tables[j].Name
	})
	return tables
}

func findLikelyTableReplacement(removed Table, newTables []Table) (replacement Table, found bool) {
	if v, ok := RenamedTables[removed.Name]; ok {
		for _, nt := range newTables {
			if nt.Name != v {
				continue
			}
			nt.Status = "renamed"
			nt.Comment = fmt.Sprintf("Renamed to [%s](#%s)", nt.Name, nt.Name)
			return nt, true
		}
	}
	plural := pluralize.NewClient()
	normalize := func(s string) string {
		return strings.ReplaceAll(s, "_", "")
	}
	singularTableName := func(name string) string {
		parts := strings.Split(name, "_")
		last := plural.Singular(parts[len(parts)-1])
		return strings.Join(append(parts[:len(parts)-1], last), "_")
	}
	normalizedName := normalize(removed.Name)
	for _, t := range newTables {
		if t.Name == removed.Name {
			continue
		}
		singular := singularTableName(t.Name)

		if normalizedName == normalize(t.Name) {
			replacement = t
			replacement.Status = "renamed"
			replacement.Comment = fmt.Sprintf("Renamed to [%s](#%s)", replacement.Name, replacement.Name)
			found = true
		} else if strings.HasPrefix(normalizedName, normalize(singular)) && (!found || len(replacement.Name) > len(t.Name)) {
			replacement = t
			replacement.Status = "moved"
			replacement.Comment = fmt.Sprintf("Moved to JSON column on [%s](#%s)", replacement.Name, replacement.Name)
			found = true
		}
	}
	return replacement, found
}

func compareColumns(t1, t2 Table) []Column {
	comparison := make(map[string]Column, 0)
	v1Map := columnsToMap(t1.Columns)
	v2Map := columnsToMap(t2.Columns)
	for name, c1 := range v1Map {
		c2, found := v2Map[name]
		if found {
			if c1.Type != c2.Type {
				c1.Status = "updated"
				c1.Comment = "Type changed from " + c1.Type + " to " + c2.Type
				c1.Type = c2.Type
			}
			comparison[name] = c1
		} else {
			c1.Status = "removed"
			comparison[name] = c1
		}
	}

	for name, c2 := range v2Map {
		_, found := v1Map[name]
		if found {
			continue
		}
		c2.Status = "added"
		comparison[name] = c2
	}

	cols := make([]Column, 0, len(comparison))
	for _, c := range comparison {
		cols = append(cols, c)
	}
	sort.Slice(cols, func(i, j int) bool {
		return cols[i].Name < cols[j].Name
	})
	return cols
}

func tablesToMap(tables []Table) map[string]Table {
	m := make(map[string]Table)
	for _, t := range tables {
		m[t.Name] = t
	}
	return m
}

func columnsToMap(columns []Column) map[string]Column {
	m := make(map[string]Column)
	for _, c := range columns {
		m[c.Name] = c
	}
	return m
}

func v2TypeToPostgres(t string) string {
	vt := convert.ValueTypeFromString(t)
	if vt == schema.TypeInvalid {
		panic("unknown type: " + t)
	}
	v, err := convert.SchemaTypeToPg(vt)
	if err != nil {
		panic(err)
	}
	return v
}

func normalizeV2Types(tables []Table) []Table {
	for i := range tables {
		for u := range tables[i].Columns {
			tables[i].Columns[u].Type = v2TypeToPostgres(tables[i].Columns[u].Type)
		}
	}
	return tables
}

func main() {
	flag.StringVar(&outputFile, "o", "", "markdown file to write results to")
	flag.StringVar(&v1TablesPath, "v1", "plugins/source/aws/docs/tables", "path to v1 table docs")
	flag.StringVar(&v2TablesPath, "v2", "plugins/source/aws/docs/tables-v2", "path to v2 table docs") // generate this using `go run main.go doc docs/tables-v2`
	flag.Parse()

	v1Tables := readTables(v1TablesPath)
	v2Tables := readTables(v2TablesPath)
	v2Tables = normalizeV2Types(v2Tables)

	comparison := compareTables(v1Tables, v2Tables)
	tpl, err := template.New("").Parse(outputTemplate)
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"Tables": comparison,
		"Date":   time.Now().Format(time.UnixDate),
	}
	var b bytes.Buffer
	if err := tpl.Execute(&b, data); err != nil {
		panic(err)
	}
	f, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(b.Bytes())
	if err != nil {
		panic(err)
	}
}
