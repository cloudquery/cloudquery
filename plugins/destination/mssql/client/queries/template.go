package queries

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed templates/*
var queriesFS embed.FS

func execTemplate(name string, data any) string {
	tpl := template.Must(template.New(name).ParseFS(queriesFS, "templates/*.sql.tpl"))
	var buf bytes.Buffer
	template.Must(tpl, tpl.Execute(&buf, data))
	return buf.String()
}
