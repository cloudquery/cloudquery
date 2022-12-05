# Policies and Compliance Frameworks

CloudQuery SQL Policies for AWS. See the [readme on GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/policies) for installation instructions.

{{- range .}}
## {{ .Name }}

### Requirements
{{ .Name }} requires the following tables to be synced before the policy is executed:

```
tables:
{{- range $table := .Tables }}
  - {{ $table }}
{{- end }}
```

### Queries
{{ .Name }} performs the following checks:
{{- range $query := .Queries }}
  - {{ $query.Title }}
{{- end }}

{{- end }}
