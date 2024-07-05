{{sanitizeID .Name}} {{sql .Type .PrimaryKey}}
{{- if .NotNull }} NOT NULL {{- end}}