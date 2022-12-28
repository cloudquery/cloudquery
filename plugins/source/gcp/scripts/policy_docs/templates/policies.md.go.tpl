# Policies and Compliance Frameworks

Policies are a set of standard SQL queries that can be run to check the security and compliance of your cloud resources against best practice frameworks.

This page documents the available CloudQuery SQL Policies for GCP. See the [readme on GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/gcp/policies) for installation instructions.

{{- range .}}
## {{ .Name }}

### Requirements
{{ .Name }} requires the following tables to be synced before the policy is executed:

```yaml
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
