package docs

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

type templateDataForSpec struct {
	PluginName string
	Tables     schema.Tables
}

func (g *Generator) renderTablesAsSpec(dir string) error {
	t, err := template.New("all_tables.yaml.go.tpl").Funcs(template.FuncMap{
		"indentToDepth": indentToDepth,
	}).ParseFS(templatesFS, "templates/all_tables.yaml.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse template for README.md: %v", err)
	}

	var b bytes.Buffer
	if err := t.Execute(&b, templateDataForSpec{PluginName: g.pluginName, Tables: g.tables}); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}
	content := b.String()
	outputPath := filepath.Join(dir, "all_tables.yaml")
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file %v: %v", outputPath, err)
	}
	defer f.Close()
	f.WriteString(content)
	return nil
}
