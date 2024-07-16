package docs

import (
	"embed"
	"fmt"
	"os"
	"regexp"
	"sort"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

var reMatchNewlines = regexp.MustCompile(`\n{3,}`)
var reMatchHeaders = regexp.MustCompile(`(#{1,6}.+)\n+`)

type Format int

const (
	FormatMarkdown Format = iota
	FormatJSON
	FormatSpec
)

func (r Format) String() string {
	return [...]string{"markdown", "json", "spec"}[r]
}

func FormatFromString(s string) (Format, error) {
	switch s {
	case "markdown":
		return FormatMarkdown, nil
	case "json":
		return FormatJSON, nil
	case "spec":
		return FormatSpec, nil
	default:
		return FormatMarkdown, fmt.Errorf("unknown format %s", s)
	}
}

type Generator struct {
	tables     schema.Tables
	pluginName string
}

func sortTables(tables schema.Tables) {
	sort.SliceStable(tables, func(i, j int) bool {
		return tables[i].Name < tables[j].Name
	})

	for _, table := range tables {
		sortTables(table.Relations)
	}
}

// NewGenerator creates a new generator for the given tables.
// The tables are sorted by name. pluginName is optional and is used in markdown only
func NewGenerator(pluginName string, tables schema.Tables) *Generator {
	sortedTables := make(schema.Tables, 0, len(tables))
	for _, t := range tables {
		sortedTables = append(sortedTables, t.Copy(nil))
	}
	sortTables(sortedTables)

	return &Generator{
		tables:     sortedTables,
		pluginName: pluginName,
	}
}

func (g *Generator) Generate(dir string, format Format) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	switch format {
	case FormatMarkdown:
		return g.renderTablesAsMarkdown(dir)
	case FormatJSON:
		return g.renderTablesAsJSON(dir)
	case FormatSpec:
		return g.renderTablesAsSpec(dir)
	default:
		return fmt.Errorf("unsupported format: %v", format)
	}
}
