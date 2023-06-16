{{sanitizeID .Name}} {{.Type | sql}}
{{- if .NotNull }} NOT NULL {{- end}}