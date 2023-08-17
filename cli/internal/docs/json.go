package docs

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

type jsonTable struct {
	Name        string       `json:"name"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Columns     []jsonColumn `json:"columns"`
	Relations   []jsonTable  `json:"relations"`
}

type jsonColumn struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	IsPrimaryKey     bool   `json:"is_primary_key,omitempty"`
	IsIncrementalKey bool   `json:"is_incremental_key,omitempty"`
}

func (g *Generator) renderTablesAsJSON(dir string) error {
	jsonTables := g.jsonifyTables(g.tables)
	buffer := &bytes.Buffer{}
	m := json.NewEncoder(buffer)
	m.SetIndent("", "  ")
	m.SetEscapeHTML(false)
	err := m.Encode(jsonTables)
	if err != nil {
		return err
	}
	outputPath := filepath.Join(dir, "__tables.json")
	return os.WriteFile(outputPath, buffer.Bytes(), 0644)
}

func (g *Generator) jsonifyTables(tables schema.Tables) []jsonTable {
	jsonTables := make([]jsonTable, len(tables))
	for i, table := range tables {
		jsonColumns := make([]jsonColumn, len(table.Columns))
		for c, col := range table.Columns {
			jsonColumns[c] = jsonColumn{
				Name:             col.Name,
				Type:             col.Type.String(),
				IsPrimaryKey:     col.PrimaryKey,
				IsIncrementalKey: col.IncrementalKey,
			}
		}
		jsonTables[i] = jsonTable{
			Name:        table.Name,
			Title:       table.Title,
			Description: table.Description,
			Columns:     jsonColumns,
			Relations:   g.jsonifyTables(table.Relations),
		}
	}
	return jsonTables
}
