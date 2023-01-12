package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

func generateTable(table *Table, dir string) error {
	var tpl *template.Template

	tpl, err := template.New("table.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower}).ParseFS(templatesFS, path.Join("templates/table.go.tpl"))
	if err != nil {
		return err
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, table); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	filePath := path.Join(dir, table.SubService+".go")

	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}

	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

func generateFetch(table *Table, dir string) error {
	var tpl *template.Template

	tpl, err := template.New("fetch.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower}).ParseFS(templatesFS, path.Join("templates/fetch.go.tpl"))
	if err != nil {
		return nil
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, table); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	filePath := path.Join(dir, table.SubService+"_fetch.go")

	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}

	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

func generateTables(tables []Table) error {
	for _, table := range tables {
		_, runtime_filename, _, ok := runtime.Caller(0)
		if !ok {
			return fmt.Errorf("failed to get caller information")
		}

		dir := path.Join(path.Dir(runtime_filename), "../resources/services/", table.Service)

		if err := generateTable(&table, dir); err != nil {
			return err
		}
	}
	return nil
}

func generateFetchers(tables []Table) error {
	for _, table := range tables {
		_, runtime_filename, _, ok := runtime.Caller(0)
		if !ok {
			return fmt.Errorf("failed to get caller information")
		}

		dir := path.Join(path.Dir(runtime_filename), "../resources/services/", table.Service)

		if err := generateFetch(&table, dir); err != nil {
			return err
		}
	}
	return nil
}

func generateTableList(tables []Table) error {
	sort.Slice(tables, func(i int, j int) bool {
		if tables[i].Service < tables[j].Service {
			return true
		}
		if tables[i].Service > tables[j].Service {
			return false
		}
		return tables[i].SubService < tables[j].SubService
	})

	tpl, err := template.New("table_list.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, "templates/table_list.go.tpl")

	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, tables); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	_, runtime_filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	filePath := path.Join(path.Dir(runtime_filename), "../resources/plugin/autogen_tables.go")
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil

}
