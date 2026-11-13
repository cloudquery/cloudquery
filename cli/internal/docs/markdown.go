package docs

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

type templateData struct {
	PluginName string
	Tables     schema.Tables
}

func (g *Generator) renderTablesAsMarkdown(dir string) error {
	for _, table := range g.tables {
		if err := g.renderAllTables(dir, table); err != nil {
			return err
		}
	}
	t, err := template.New("all_tables.md.go.tpl").Funcs(template.FuncMap{
		"indentToDepth": indentToDepth,
	}).ParseFS(templatesFS, "templates/all_tables*.md.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse template for README.md: %v", err)
	}

	var b bytes.Buffer
	if err := t.Execute(&b, templateData{PluginName: g.pluginName, Tables: g.tables}); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}
	content := formatMarkdown(b.String())
	outputPath := filepath.Join(dir, "README.md")
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file %v: %v", outputPath, err)
	}
	defer f.Close()
	f.WriteString(content)
	return nil
}

func (g *Generator) renderAllTables(dir string, t *schema.Table) error {
	if err := g.renderTable(dir, t); err != nil {
		return err
	}
	for _, r := range t.Relations {
		if err := g.renderAllTables(dir, r); err != nil {
			return err
		}
	}
	return nil
}

func MarkdownFilePath(dir string, name string) (string, error) {
	fileName := name + ".md"
	if fileName != filepath.Base(fileName) || !filepath.IsLocal(fileName) {
		return "", fmt.Errorf("invalid name %q", name)
	}
	return filepath.Join(dir, fileName), nil
}

func (*Generator) renderTable(dir string, table *schema.Table) error {
	t := template.New("")
	t, err := t.New("table.md.go.tpl").ParseFS(templatesFS, "templates/table.md.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	outputPath, err := MarkdownFilePath(dir, table.Name)
	if err != nil {
		return err
	}

	var b bytes.Buffer
	if err := t.Execute(&b, table); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}
	content := formatMarkdown(b.String())
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file %v: %v", outputPath, err)
	}
	defer f.Close()
	f.WriteString(content)
	return f.Close()
}

func formatMarkdown(s string) string {
	s = reMatchNewlines.ReplaceAllString(s, "\n\n")
	return reMatchHeaders.ReplaceAllString(s, `$1`+"\n\n")
}

func indentToDepth(table *schema.Table) string {
	s := ""
	t := table
	for t.Parent != nil {
		s += "  "
		t = t.Parent
	}
	return s
}
